package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Create creates and returns handler object
// that is concerned with routing of server request paths
func createRouter() *mux.Router {
	// Router
	router := mux.NewRouter()
	// REST api serving
	apirouter := router.PathPrefix("/api/notes").Subrouter()
	apirouter.HandleFunc("/", getNotes).Methods(http.MethodGet)
	apirouter.HandleFunc("/{id:[0-9]+}", getNote).Methods(http.MethodGet)
	apirouter.HandleFunc("/", postNote).Methods(http.MethodPost)
	apirouter.HandleFunc("/{id:[0-9]+}", deleteNote).Methods(http.MethodDelete)
	// SPA serving
	spa := spaHandler{staticPath: "web", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	return router
}
