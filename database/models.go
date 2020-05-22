package database

import "time"

// NotesResource is the model for holding notes information
type notesResource struct {
	ID    int
	Stamp time.Time
	Note  string
	Tags  string
}

// Model to ccreate notes database
const createNotesSQL = `CREATE TABLE IF NOT EXISTS notes (
						ID INTEGER PRIMARY KEY AUTOINCREMENT,
						STAMP DATE,
						NOTE TEXT NOT NULL,
						TAGS VARCHAR(255)
						) `
