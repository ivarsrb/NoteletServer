package main

import (
	"github.com/ivarsrb/NoteletServer/database"
	"github.com/ivarsrb/NoteletServer/server"

	// For hosting on Heroku
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	database.Create()
	server.Run()
}
