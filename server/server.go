/*
Package server implements setting up web server, defines routing and handles protocol requests.
Server configuration is read from command line parameters upon initialization.
*/
package server

import (
	"flag"
	"log"
	"net/http"
	"time"
)

// config stores server confguration settings
type config struct {
	// Network address
	addr string
	// Path to static web files
	staticPath string
}

var cfg config

func init() {
	// Define configuration flags, used to configure the system upon execution
	flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.staticPath, "static", "web", "Path to static web files")
	flag.Parse()
}

// Serve creates and listne for requests
func Serve() {
	router := createRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         cfg.addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Printf("Serving on '%s' ...", cfg.addr)
	log.Fatal(server.ListenAndServe())
}
