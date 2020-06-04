package main

import (
	"github.com/ivarsrb/NoteletServer/database"
	"github.com/ivarsrb/NoteletServer/server"
)

func main() {
	database.Create()
	server.Run()
}
