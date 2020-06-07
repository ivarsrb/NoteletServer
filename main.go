package main

import (
	"log"

	"github.com/ivarsrb/NoteletServer/server"
	"github.com/ivarsrb/NoteletServer/storage"

	// For hosting on Heroku
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	var err error
	// TODO maybe return storage object here and pass on rather
	// than store it clobaly in storage package
	err = storage.New(storage.SQLite)
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
}
