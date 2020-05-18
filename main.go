package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// --------------------------------
// REST Api
// api/notes/  GET Get the list of all available notes
// api/notes/(id)  GET Get a particular note
// api/notes/ POST Add new note
// api/notes/(id) PUT Modifie given note (by overwriting a record)
// api/notes/(id) DELETE Delete a particular note

func main() {
	router := mux.NewRouter()
	// REST api serving
	subrouter := router.PathPrefix("/api/notes").Subrouter()
	subrouter.HandleFunc("/", notesHandler)
	subrouter.HandleFunc("/{id:[0-9]+}", noteHandler)
	// SPA serving
	spa := spaHandler{staticPath: "templates", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
