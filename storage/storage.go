package storage

import "github.com/ivarsrb/NoteletServer/notes"

// Type defines available storage types
type Type int

const (
	// SQLite stores data in SQLite database
	SQLite Type = iota
	// PostgresSQL stores data in PostgresSQL database
	PostgresSQL
)

// Storage is an interface for possible databases that could be used.
// They all should implement this interface.
type Storage interface {
	InsertNote() error
	SelectNote(int) (notes.Note, error)
	SelectNotes() ([]notes.Note, error)
	DeleteNote(int) error
}

// DB is the global storage instance
var DB Storage

// New instantiates global storage of the given type
func New(t Type) error {
	var err error
	//var err error
	switch t {
	case SQLite:
		//DB, err = NewSQLite("./notelet.db")
		if err != nil {
			return err
		}
	case PostgresSQL:

	}

	return nil
}
