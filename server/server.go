/*
Package server implements setup of the web server, defines routing and handles protocol requests.
Server configuration is read from env variables
*/
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ivarsrb/NoteletServer/logger"
)

// Run creates and and starts http server on the given port
func Run(port string) error {
	router := newRouter()
	timeout := 10 * time.Second
	maxHeaderSize := 1 << 20
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
		MaxHeaderBytes: maxHeaderSize,
		ErrorLog:       logger.Error,
	}
	// Initializing the server in a goroutine because graceful shutdown
	// is listening for interrupt signals from OS
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatal("server: error while listening and serving: ", err)
		}
	}()
	var err error
	err = gracefulShutdown(server)
	if err != nil {
		return fmt.Errorf("server shutdown error:: %v", err)
	}
	logger.Info.Println("server: exiting")
	return nil
}

// gracefulShutdown listens for signals if server is interrupted and tries
// to shutdown with little damage as possible
func gracefulShutdown(server *http.Server) error {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info.Println("server: shutting down...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
