/*
Package storage provides data storage manipulation by implementing Storage inerface
by desired database type
*/
package storage

import (
	"fmt"

	"github.com/ivarsrb/NoteletServer/notes"
)

// Type defines available storage types
type Type int

const (
	// PostgresSQL specifies the database type
	PostgresSQL Type = iota
	// Add new storege types here .. and in New() function
)

// Storage is an interface for possible database types that could be used.
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

// New instantiates global storage of the given type with given url string
// that is interpreted depending on storage type
func New(t Type, dburl string) error {
	var err error
	switch t {
	case PostgresSQL:
		DB, err = NewPostgres(dburl)
		if err != nil {
			return fmt.Errorf("data storage creation fail: %v", err)
		}
	}

	return nil
}
