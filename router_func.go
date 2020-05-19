package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Middlawae function
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

// Notes request handling
func notesHandler(w http.ResponseWriter, r *http.Request) {

	// Middleware

	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"notes list": true})

}

// Notes request handling
func noteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// an example API handler
	json.NewEncoder(w).Encode(map[string]bool{"individual note " + vars["id"]: true})
}
