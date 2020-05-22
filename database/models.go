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
