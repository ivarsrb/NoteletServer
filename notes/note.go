package notes

import "time"

// Note stores information about the note
type Note struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
	Note      string    `json:"note" binding:"required"`
	Tags      string    `json:"tags"`
}
