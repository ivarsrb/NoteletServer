package main

import (
	"fmt"
	"log"
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
func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorization screen: %s", r.URL.Path)
}

// Notes request handling
func notesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Notes management: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", authorizationHandler)
	http.HandleFunc("/notes/", notesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
