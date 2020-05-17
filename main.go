package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// --------------------------------
// REST Api
// /notes/  GET Get the list of all available notes
// /notes/(id)  GET Get a particular note
// /notes/ POST Add new note
// /notes/(id) PUT Modifie given note (by overwriting a record)
// /notes/(id) DELETE Delete a particular note

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/notes", notesHandler)
	router.GET("/notes/:note", notesHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
