/*
Package app ties together the application and holds it's state
*/
package app

import (
	"fmt"
	"os"

	"github.com/ivarsrb/NoteletServer/server"
	"github.com/ivarsrb/NoteletServer/storage"
)

// Application structure stores and application state
type Application struct {
	// port stores address from where the server is accessible
	port string
	// dBurl stores database access string
	dBurl string
}

// Run method initializes application objects and starts execution
// If error occured in any of those objects, execution will be stopped and error returned
func (a *Application) Run() error {
	var err error
	err = a.initParams()
	if err != nil {
		return fmt.Errorf("unable to init params: %v", err)
	}
	err = storage.New(storage.PostgresSQL, a.dBurl)
	if err != nil {
		return err
	}
	err = server.Run(a.port)
	if err != nil {
		return err
	}

	return nil
}

// initParams initializes state variables, return error if unable
func (a *Application) initParams() error {
	// Gather environemnt variables
	a.port = os.Getenv("PORT")
	if a.port == "" {
		return fmt.Errorf("$PORT must be set")
	}
	a.dBurl = os.Getenv("DATABASE_URL")
	if a.dBurl == "" {
		return fmt.Errorf("$DATABASE_URL must be set")
	}
	return nil
}
