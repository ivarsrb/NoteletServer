package main

import (
	"fmt"
	"net/http"
)

// --------------------------------
// REST Api
// /notes/  GET Get the list of all available notes
// /notes/(id)  GET Get a particular note
// /notes/ POST Add new note
// /notes/(id) PUT Modifie given note (by overwriting a record)
// /notes/(id) DELETE Delete a particular note

// TODO: need multiplexer

// Authorization screen (login)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index screen: %s", r.URL.Path)
}

// Notes request handling
func notesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Notes management: %s", r.URL.Path)
}

func main() {
	// Multiplexer routes requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/notes/", notesHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
