package storage

import (
	"database/sql"

	"github.com/ivarsrb/NoteletServer/notes"
	// To initialize sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// createNotesSQL stores SQL script to create database
const createNotesSQL = `CREATE TABLE IF NOT EXISTS notes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp DATE,
	note TEXT NOT NULL,
	tags VARCHAR(255)
	) `

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
		//logger.Error.Fatal(err)
		return nil, err
	}
	//logger.Info.Println("Database: first connection successfully established")
	return &stg, nil
}

// createDatabase creates the database with the model script provided
func createDatabase(db *sql.DB) error {
	// Create train table
	_, err := db.Exec(createNotesSQL)
	if err != nil {
		return err
	}
	return nil
}

// SelectNotes is ....
func (s *SQLiteStorage) SelectNotes() ([]notes.Note, error) {
	return nil, nil
}
