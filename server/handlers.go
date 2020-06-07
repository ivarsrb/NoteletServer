package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivarsrb/NoteletServer/database"
	"github.com/ivarsrb/NoteletServer/logger"

	"github.com/ivarsrb/NoteletServer/storage"
)

// GetNotes retrieve a list of all notes
func GetNotes(c *gin.Context) {
	notes, err := storage.DB.SelectNotes()
	if err != nil {
		logger.Error.Println("Server: error retrieving notes!", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot retrieve note list"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// GetNote retrieve a note with a requested id from the database
// and send it back as a json
func GetNote(c *gin.Context) {
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
	var note database.NoteResource
	if note.Get(id) {
		c.JSON(http.StatusOK, note)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// PostNote adds new note
func PostNote(c *gin.Context) {
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
	var note database.NoteResource
	if err := c.ShouldBindJSON(&note); err != nil {
		logger.Error.Println("Server:" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Add()
	//w.Header().Set("Content-Type", "text/plain")
	//w.WriteHeader(http.StatusCreated)
	c.Status(http.StatusCreated)
}

// DeleteNote delete a given note
func DeleteNote(c *gin.Context) {
	var id int
	var err error
	// Identifier which note to delete.
	// It should be integer
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		logger.Error.Println("Server: 'id' should be an integer type")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID of an unsupported type!"})
		return
	}
	//w.Header().Set("Content-Type", "text/plain")
	// Check if there was something to delete
	if database.DeleteNote(id) {
		c.Status(http.StatusOK)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
