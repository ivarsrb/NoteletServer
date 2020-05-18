package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Notes request handling
func notesHandler(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"notes list": true})
}

// Notes request handling
func noteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"individual note " + vars["id"]: true})
}
