package main

import (
	// For hosting on Heroku
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/ivarsrb/NoteletServer/app"
	"github.com/ivarsrb/NoteletServer/logger"
)

func main() {
	var app app.Application
	err := app.Run()
	if err != nil {
		logger.Error.Fatal(err)
	}
}
