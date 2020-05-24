package database

import (
	"log"
	"time"
)

// Model to ccreate notes database
const createNotesSQL = `CREATE TABLE IF NOT EXISTS notes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp DATE,
	note TEXT NOT NULL,
	tags VARCHAR(255)
	) `

// NoteResource is the model for holding notes information
type NoteResource struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Note      string    `json:"note"`
	Tags      string    `json:"tags"`
}

// Get retrieves note with the given id from the database
// and writes to reciever struct.
// Returns true if record exists, false if does not exist
func (n *NoteResource) Get(id int) bool {
	err := db.QueryRow("SELECT id, timestamp, note, tags FROM notes where id = ?", id).
		Scan(&n.ID, &n.Timestamp, &n.Note, &n.Tags)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Add adds this note to database and writes ID from just inserted record
func (n *NoteResource) Add() {
	stmt, err := db.Prepare("INSERT INTO notes(id, timestamp, note, tags ) VALUES (NULL,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(n.Timestamp, n.Note, n.Tags)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	n.ID = int(lastID)
}

// GetNotesAll retrieves all notes from the database
// and return as a notes resource slice
func GetNotesAll() []NoteResource {
	// Make an empty slice. Null slice like "var notes []NoteResource" will
	// json marshal into 'null'.
	notes := make([]NoteResource, 0)
	//notes = append(notes, NoteResource{})
	rows, err := db.Query("select id, timestamp, note, tags FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	// Rows should be closed to avoid connection holding
	defer rows.Close()
	// Iterate over all selected rows and append to return slice
	for rows.Next() {
		var note NoteResource
		err := rows.Scan(&note.ID, &note.Timestamp, &note.Note, &note.Tags)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	// Check for abdnormal loop termination
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return notes
}

// DeleteNote removes note with the given id from the database
// Returns true if record existed before deletion, false - if did not
func DeleteNote(id int) bool {
	stmt, err := db.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	// Note: not every driver may support this feature
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowCnt != 1 {
		return false
	}
	return true
}
