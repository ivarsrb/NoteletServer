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

	"github.com/ivarsrb/NoteletServer/logger"

	// For hosting on Heroku
	_ "github.com/heroku/x/hmetrics/onload"
)

// Run creates and and starts http server
func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		logger.Error.Fatal("$PORT must be set")
	}

	router := newRouter()
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Initializing the server in a goroutine because graceful shutdown
	// is listening for interrupt signals from OS
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("listen: %s\n", err)
		}
	}()
	gracefulShutdown(server)
	logger.Info.Println("Server exiting")
}

// gracefulShutdown listens for signals if server is interrupted and tries
// to shutdown with little damage as possible
func gracefulShutdown(server *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
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
}
