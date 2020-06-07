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
		logger.Error.Println("Server: error retrieving notes!", err)
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
		logger.Error.Println("Server: 'id' should be an integer type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID of an unsupported type"})
		return
	}
	// Retrieve the record if possible
	note, err := storage.DB.SelectNote(id)
	if err != nil {
		logger.Error.Println("Server: error retrieving a note.", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to retrieve a record with the given id"})
	} else {
		c.JSON(http.StatusOK, note)
	}
}

// postNote adds new note
func postNote(c *gin.Context) {
	var err error
	// Check for appropriate content type
	contentType := c.Request.Header.Get("Content-type")
	if contentType != "application/json" {
		logger.Error.Println("Server: request content type is not 'application/json'")
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "JSON content type expected!"})
		return
	}
	/*
		// Set limit for maximum body size
		// A request body larger will result in
		// Decode() returning a "http: request body too large" error.
		const maxBodySize = 1048576
		r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
		// Decode the json data we recieved
		decoder := json.NewDecoder(r.Body)
		// Dissalow any unsupported fields incoming from request
		decoder.DisallowUnknownFields()
		var note database.NoteResource
		err := decoder.Decode(&note)
		if err != nil {
			logger.Error.Println(err)
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			return
		}*/
	var note notes.Note
	if err = c.ShouldBindJSON(&note); err != nil {
		logger.Error.Println("Server: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse recieved json"})
		return
	}
	err = storage.DB.InsertNote(&note)
	if err != nil {
		logger.Error.Println("Server: error inserting a note .", err)
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
		logger.Error.Println("Server: 'id' should be an integer type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID of an unsupported type!"})
		return
	}
	err = storage.DB.DeleteNote(id)
	if err != nil {
		logger.Error.Println("Server: error deleting a note.", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to delete a record with the given id"})
	} else {
		c.Status(http.StatusOK)
	}
}
