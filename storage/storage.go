/*
Package storage package handles storage initialization and implement storage interface
*/
package storage

import (
	"fmt"

	"github.com/ivarsrb/NoteletServer/notes"
)

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
	InsertNote(note *notes.Note) error
	SelectNote(int) (notes.Note, error)
	SelectNotes(filter string) ([]notes.Note, error)
	DeleteNote(int) error
	UpdateNote(note *notes.Note) error
}

// DB is the global storage instance
var DB Storage

// New instantiates global storage of the given type
func New(t Type) error {
	var err error
	//var err error
	switch t {
	case SQLite:
		DB, err = NewSQLite("./notelet.db")
		if err != nil {
			return fmt.Errorf("data storage creation fail: %v", err)
		}
	case PostgresSQL:
		DB, err = NewPostgres("postgres://testusr:testpass123@localhost/testdb?sslmode=disable")
		if err != nil {
			return fmt.Errorf("data storage creation fail: %v", err)
		}
	}

	return nil
}
