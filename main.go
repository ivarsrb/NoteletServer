package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ivarsrb/NoteletServer/database"
)

// --------------------------------
// REST Api
// api/notes/  GET Get the list of all available notes
// api/notes/(id)  GET Get a particular note
// api/notes/ POST Add new note
// api/notes/(id) PUT Modify given note (by overwriting a record)
// api/notes/(id) DELETE Delete a particular note

func main() {
	// Initialize database (create tables if necessery)
	database.Initialize()
	// Router
	router := mux.NewRouter()
	// REST api serving
	subrouter := router.PathPrefix("/api/notes").Subrouter()
	subrouter.HandleFunc("/", notesHandler)
	subrouter.HandleFunc("/{id:[0-9]+}", noteHandler)
	// SPA serving
	spa := spaHandler{staticPath: "tegmplates", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("Serving ...")
	log.Fatal(server.ListenAndServe())
}
