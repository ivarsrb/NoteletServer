package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ivarsrb/NoteletServer/database"
)

// getNotes retrieve a list of all notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	//json.NewEncoder(w).Encode(map[string]bool{"notes list": true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get notes called"}`))
}

// getNote retrieve a note with a requested id from the database
// and send it back as a json
func getNote(w http.ResponseWriter, r *http.Request) {
	var id int
	var err error
	// Identifier which note to get.
	// It should be integer (it is also checked at router with regexp)
	if id, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		log.Println("Handlers: 'id' should be an integer type")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
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

// postNote add new note
func postNote(w http.ResponseWriter, r *http.Request) {
	// An example API handler
	//json.NewEncoder(w).Encode(map[string]bool{"notes list": true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "create note called"}`))
}

// deleteNote delete a given note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// an example API handler
	//json.NewEncoder(w).Encode(map[string]bool{"notes list": true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message": "delete note called, id %s"}`, vars["id"])))
}
