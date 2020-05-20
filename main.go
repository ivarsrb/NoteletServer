package main

import (
	"github.com/ivarsrb/NoteletServer/database"
	"github.com/ivarsrb/NoteletServer/server"
)

func main() {
	// Initialize database (create tables if necessery)
	database.Initialize()
	server.Serve()
}
