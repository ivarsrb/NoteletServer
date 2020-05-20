package server

import (
	"log"
	"net/http"
	"time"
)

// Serve creates and listne for requests
func Serve() {
	router := createRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("Serving ...")
	log.Fatal(server.ListenAndServe())
}
