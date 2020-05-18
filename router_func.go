package main

import (
	"encoding/json"
	"net/http"
)

// Notes request handling
func notesHandler(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
