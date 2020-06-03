/*
Package server implements setting up web server, defines routing and handles protocol requests.
Server configuration is read from command line parameters upon initialization.
*/
package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ivarsrb/NoteletServer/logger"

	// For hosting on Heroku
	_ "github.com/heroku/x/hmetrics/onload"
)

// Serve creates and listne for requests
func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		logger.Error.Fatal("$PORT must be set")
	}

	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Serve frontend static files
	//router.Use(static.Serve("/", static.LocalFile("./web", true)))
	// This may have problems when we need js files served, look in github help for other examples
	router.StaticFile("/", "./web/index.html")

	// Setup route group for the API
	api := router.Group("/api")
	// Authorization middleware
	// api.Use(AuthRequired())
	api.GET("/notes", getNotes)
	api.GET("/notes/:id", getNote)
	api.POST("/notes", postNote)
	api.DELETE("/notes/:id", deleteNote)
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Graceful shutdown
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info.Println("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error.Fatal("Server forced to shutdown:", err)
	}

	logger.Info.Println("Server exiting")
}
