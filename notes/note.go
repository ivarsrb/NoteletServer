/*
Package notes is responsible for basic knowledge of the note type
*/
package notes

import "time"

// Note stores information about the note
type Note struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Note      string    `json:"note" binding:"required"`
	Tags      string    `json:"tags"`
}
