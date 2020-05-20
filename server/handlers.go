package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// getNotes retrieve a list of all notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	//json.NewEncoder(w).Encode(map[string]bool{"notes list": true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get notes called"}`))
}

// getNote retrieve a notw with a particular id
func getNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// an example API handler
	//json.NewEncoder(w).Encode(map[string]bool{"individual note " + vars["id"]: true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message": "get note called, id %s"}`, vars["id"])))
}

// postNote add new note
func postNote(w http.ResponseWriter, r *http.Request) {
	// an example API handler
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
