package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivarsrb/NoteletServer/logger"
	"github.com/ivarsrb/NoteletServer/notes"

	"github.com/ivarsrb/NoteletServer/storage"
)

// getNotes retrieve a list of all notes
func getNotes(c *gin.Context) {
	notes, err := storage.DB.SelectNotes()
	if err != nil {
		logger.Error.Println("server: error retrieving notes: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot retrieve note list"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// getNote retrieve a note with a requested id from the database and send it back as a json
func getNote(c *gin.Context) {
	var id int
	var err error
	// Identifier which note to get.
	// It should be integer
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		logger.Error.Println("server: 'id' should be an integer type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID of an unsupported type"})
		return
	}
	// Retrieve the record if possible
	note, err := storage.DB.SelectNote(id)
	if err != nil {
		logger.Error.Println("server: error retrieving a note: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to retrieve a record with the given id"})
		return
	}
	c.JSON(http.StatusOK, note)
}

// postNote adds new note
func postNote(c *gin.Context) {
	var err error
	// Check for appropriate content type
	contentType := c.Request.Header.Get("Content-type")
	if contentType != "application/json" {
		logger.Error.Println("server: request content type is not 'application/json'")
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "JSON content type expected!"})
		return
	}

	var note notes.Note
	if err = c.ShouldBindJSON(&note); err != nil {
		logger.Error.Println("server: cannot parse note json recieved from client: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse recieved json"})
		return
	}
	err = storage.DB.InsertNote(&note)
	if err != nil {
		logger.Error.Println("server: error inserting a note: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add the not"})
		return
	}
	c.Status(http.StatusCreated)
}

// deleteNote delete a given note
func deleteNote(c *gin.Context) {
	var id int
	var err error
	// Identifier which note to delete.
	// It should be integer
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		logger.Error.Println("server: 'id' should be an integer type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID of an unsupported type!"})
		return
	}
	err = storage.DB.DeleteNote(id)
	if err != nil {
		logger.Error.Println("server: error deleting a note: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to delete a record with the given id"})
		return
	}
	c.Status(http.StatusOK)
}
