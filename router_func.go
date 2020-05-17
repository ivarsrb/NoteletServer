package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Authorization screen (login)
func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Index screen: %s", r.URL.Path)
}

// Notes request handling
func notesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Notes management: %s %s", r.URL.Path, ps.ByName("note"))
}
