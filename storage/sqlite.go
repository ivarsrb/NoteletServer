package storage

import (
	"database/sql"
	"fmt"

	"github.com/ivarsrb/NoteletServer/notes"

	// To initialize sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStorage implements the storage interface agains SQLite db
type SQLiteStorage struct {
	// Pointer to database connection pool
	db *sql.DB
}

// NewSQLite creates and returns SQLite storage object.
// The database connection is established and database created.
// name parameter specifies database file name
func NewSQLite(name string) (*SQLiteStorage, error) {
	var err error
	stg := SQLiteStorage{}
	// Prepares the database abstraction for later use
	// The first actual connection to the underlying datastore will
	// be established lazily, when it’s needed for the first time.
	const driver = "sqlite3"
	stg.db, err = sql.Open(driver, name)
	if err != nil {
		//logger.Error.Fatal(err)
		return nil, err
	}
	// Actual first connection with the database is established here.
	// Create storage table by executing a script.
	err = createDatabase(stg.db)
	if err != nil {
		return nil, fmt.Errorf("database '%s' creation fail: %v", name, err)
	}
	//logger.Info.Println("Database: first connection successfully established")
	return &stg, nil
}

// createDatabase creates the database with the model script provided
func createDatabase(db *sql.DB) error {
	// createNotesSQL stores SQL script to create database
	const createNotesSQL = `CREATE TABLE IF NOT EXISTS notes (
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								timestamp DATE,
								note TEXT NOT NULL,
								tags VARCHAR(255)
								) `
	// Create train table
	_, err := db.Exec(createNotesSQL)
	if err != nil {
		return err
	}
	return nil
}

// SelectNotes retrieves all notes from the database
// and return as a notes resource slice
func (s *SQLiteStorage) SelectNotes() ([]notes.Note, error) {
	// Make an empty slice. Null slice like "var notes []NoteResource" will
	// json marshal into 'null'.
	noteList := make([]notes.Note, 0)
	rows, err := s.db.Query("SELECT id, timestamp, note, tags FROM notes")
	if err != nil {
		return nil, err
	}
	// Rows should be closed to avoid connection holding
	defer rows.Close()
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
func (s *SQLiteStorage) SelectNote(id int) (notes.Note, error) {
	var note notes.Note
	err := s.db.QueryRow("SELECT id, timestamp, note, tags FROM notes where id = ?", id).
		Scan(&note.ID, &note.Timestamp, &note.Note, &note.Tags)
	if err != nil {
		return notes.Note{}, fmt.Errorf("note with id '%d': %v", id, err)
	}
	return note, nil
}

// InsertNote adds given note to database
func (s *SQLiteStorage) InsertNote(note *notes.Note) error {
	var err error
	stmt, err := s.db.Prepare("INSERT INTO notes(id, timestamp, note, tags) VALUES (NULL,?,?,?)")
	if err != nil {
		return err
	}
	/*res*/ _, err = stmt.Exec(note.Timestamp, note.Note, note.Tags)
	if err != nil {
		return err
	}
	/*
		lastID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		n.ID = int(lastID)
	*/
	return nil
}

// DeleteNote removes note with the given id from the database
func (s *SQLiteStorage) DeleteNote(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM notes WHERE id = ?")
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
func (s *SQLiteStorage) UpdateNote(note *notes.Note) error {
	// Update time stamp upon updating the record
	stmt, err := s.db.Prepare("UPDATE notes SET timestamp = CURRENT_TIMESTAMP, note = ?, tags = ? WHERE id = ?")
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
