package storage

import (
	// To initialize PostgreSQL driver
	"database/sql"
	"fmt"

	"github.com/ivarsrb/NoteletServer/notes"

	// To initialize sqlite3 driver
	_ "github.com/lib/pq"
)

// PostgresStorage implements the storage interface agains PostgreSQL db
type PostgresStorage struct {
	// Pointer to database connection pool
	db *sql.DB
}

// prepANdExec prepares statement and executes it in a single go.
// In case any error occures it is returned
func prepAndExec(db *sql.DB, query string) error {
	// Create index on searchable field
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// createPostgresDB creates and sets up the database
func createPostgresDB(db *sql.DB) error {
	//var stmt *sql.Stmt
	var err error
	// Create table
	err = prepAndExec(db, `CREATE TABLE IF NOT EXISTS notes (
							id SERIAL PRIMARY KEY,
							timestamp TIMESTAMP without time zone DEFAULT timezone('UTC', now()),
							note TEXT NOT NULL CONSTRAINT notechk CHECK (char_length(note) <= 10000),
							tags VARCHAR(255),
							tsv TSVECTOR
							)`)
	if err != nil {
		return err
	}
	// Create index on searchable field
	err = prepAndExec(db, `CREATE INDEX IF NOT EXISTS notes_idx ON notes USING GIN(tsv)`)
	if err != nil {
		return err
	}
	// Function for storing text-search tokens extracted from notes and tags columns
	// with assigned priorities
	// Function to_tsvector() take argument of a dictionary to use when normalizing the words
	// as lexemes
	err = prepAndExec(db, `CREATE OR REPLACE FUNCTION search_trigger() RETURNS trigger AS $$
							BEGIN
								NEW.tsv :=
									setweight(to_tsvector(NEW.tags), 'A') ||
									setweight(to_tsvector(NEW.note), 'B');
								return NEW;
							END
							$$ LANGUAGE plpgsql;`)
	if err != nil {
		return err
	}
	// Drop trigger if it exists.
	// There is no mechanism to do it in one statement.
	err = prepAndExec(db, `DROP TRIGGER IF EXISTS tsvector_update ON notes`)
	if err != nil {
		return err
	}
	// When new record is added or updated call the function that extracts
	// text search triggers and stores them
	err = prepAndExec(db, `CREATE TRIGGER tsvector_update BEFORE INSERT OR UPDATE
							ON notes 
							FOR EACH ROW EXECUTE PROCEDURE search_trigger()`)
	if err != nil {
		return err
	}

	return nil
}

// NewPostgres creates and returns NewPostgreSQL storage object.
// The database connection is established and database created.
// name parameter specifies database file name
func NewPostgres(name string) (*PostgresStorage, error) {
	var err error
	stg := PostgresStorage{}
	// Prepares the database abstraction for later use
	// The first actual connection to the underlying datastore will
	// be established lazily, when it’s needed for the first time.
	const driver = "postgres"
	stg.db, err = sql.Open(driver, name)
	if err != nil {
		return nil, fmt.Errorf("PostgreSQL db open fail: %v", err)
	}
	// Actual first connection with the database is established here.
	// Create storage table by executing a script.
	err = createPostgresDB(stg.db)
	if err != nil {
		return nil, fmt.Errorf("PostgreSQL db creation fail: %v", err)
	}

	return &stg, nil
}

// SelectNotes retrieves filtered notes from the database
// and return as a notes resource slice
// If filter parameter is empty all notes are returned in descending order
// Filter phrase is searched in tags and note fields that are tokenized in tsv field.
func (s *PostgresStorage) SelectNotes(filter string) ([]notes.Note, error) {
	var rows *sql.Rows
	var err error
	if filter == "" {
		rows, err = s.db.Query("SELECT id, timestamp, note, tags FROM notes ORDER BY id DESC")
	} else {
		// plainto_tsquery takes plane string and tokenizes and adds & between words.
		// For more fine-tuned search use to_tquery() with custom logical operators.
		// Ranking is sorted taking weights into account (tags get more weight than note text)
		rows, err = s.db.Query(`SELECT id, timestamp, note, tags FROM notes 
								WHERE tsv @@ plainto_tsquery($1)
								ORDER BY ts_rank_cd(tsv, plainto_tsquery($1)) DESC`, filter)
	}
	if err != nil {
		return nil, err
	}
	// Rows should be closed to avoid connection holding
	defer rows.Close()
	// Make an empty slice. Null slice like "var notes []NoteResource" will
	// json marshal into 'null'.
	noteList := make([]notes.Note, 0)
	// Iterate over all selected rows and append to return slice
	for rows.Next() {
		var note notes.Note
		err := rows.Scan(&note.ID, &note.Timestamp, &note.Note, &note.Tags)
		if err != nil {
			return nil, err
		}
		noteList = append(noteList, note)
	}
	// Check for abdnormal loop termination
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return noteList, nil
}

// SelectNote retrieves and returns a note with the given id from the database
func (s *PostgresStorage) SelectNote(id int) (notes.Note, error) {
	var note notes.Note
	err := s.db.QueryRow("SELECT id, timestamp, note, tags FROM notes where id = $1", id).
		Scan(&note.ID, &note.Timestamp, &note.Note, &note.Tags)
	if err != nil {
		return notes.Note{}, fmt.Errorf("note with id '%d': %v", id, err)
	}
	return note, nil
}

// InsertNote adds given note to database
// Note body and tags are inserted from client, timestamp is set automatically upon
// record creation
func (s *PostgresStorage) InsertNote(note *notes.Note) error {
	var err error
	stmt, err := s.db.Prepare("INSERT INTO notes(note, tags) VALUES ($1,$2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(note.Note, note.Tags)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNote removes note with the given id from the database
func (s *PostgresStorage) DeleteNote(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM notes WHERE id = $1")
	if err != nil {
		return fmt.Errorf("note with id '%d': %v", id, err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("note with id '%d': %v", id, err)
	}
	// NOTE: not every driver may support this feature
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("note with id '%d': %v", id, err)
	}
	if rowCnt != 1 {
		return fmt.Errorf("numer of affected rows is not '1'")
	}
	return nil
}

// UpdateNote update a note (ID from structure) with the new
// values from the structure
// Note body and tags are updated from client, timestamp is updated on a server request
func (s *PostgresStorage) UpdateNote(note *notes.Note) error {
	stmt, err := s.db.Prepare("UPDATE notes SET timestamp = timezone('UTC', now()), note = $1, tags = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(note.Note, note.Tags, note.ID)
	if err != nil {
		return err
	}
	// NOTE: not every driver may support this feature
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("note with id '%d': %v", note.ID, err)
	}
	if rowCnt != 1 {
		return fmt.Errorf("numer of affected rows is not '1'. Id '%d'", note.ID)
	}
	return nil
}
