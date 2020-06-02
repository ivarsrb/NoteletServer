/*
Package server implements setting up web server, defines routing and handles protocol requests.
Server configuration is read from command line parameters upon initialization.
*/
package server

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	// For hosting on Heroku
	_ "github.com/heroku/x/hmetrics/onload"
)

// Serve creates and listne for requests
func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	// Setup route group for the API
	api := router.Group("/api")
	api.GET("/notes", getNotes)
	api.GET("/notes/:id", getNote)
	api.POST("/notes", postNote)
	api.DELETE("/notes/:id", deleteNote)
	router.Run(":" + port)
	/*
		router := createRouter()
		server := &http.Server{
			Handler:      router,
			Addr:         cfg.addr,
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
			ErrorLog:     logger.Error,
		}
		logger.Info.Printf("Serving on '%s' ...", cfg.addr)
		logger.Error.Fatal(server.ListenAndServe())
	*/
}
