package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ivarsrb/NoteletServer/database"
)

// getNotes retrieve a list of all notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	notes := database.GetNotesAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

// getNote retrieve a note with a requested id from the database
// and send it back as a json
func getNote(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error
	// Identifier which note to get.
	// It should be integer (it is also checked at router with regexp)
	if id, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		log.Println("Server: 'id' should be an integer type")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "ID of an unsupported type!", http.StatusBadRequest)
		return
	}
	// Retrieve the record if possible
	var note database.NoteResource
	if note.Get(id) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(note)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
	}
}

// postNote adds new note
func postNote(w http.ResponseWriter, r *http.Request) {
	// Check for appropriate content type
	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" {
		log.Println("Server: request content type is not 'application/json'")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "JSON content type expected!", http.StatusUnsupportedMediaType)
		return
	}
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
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	note.Add()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
}

// deleteNote delete a given note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error
	// Identifier which note to delete.
	// It should be integer (it is also checked at router with regexp)
	if id, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		log.Println("Handlers: 'id' should be an integer type")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "ID of an unsupported type!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	// Check if there was something to delete
	if database.DeleteNote(id) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
