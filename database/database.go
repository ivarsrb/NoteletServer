/*
Package database implements database creation and manipulation
routines.
Database connection pool and first connection are established upon initialization.
*/
package database

import (
	"database/sql"
	"log"

	// To initialize sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName   = "./notelet.db"
	dbDriver = "sqlite3"
)

// Pointer to database connection pool
var db *sql.DB

func init() {
	// Prepares the database abstraction for later use
	// The first actual connection to the underlying datastore will
	// be established lazily, when it’s needed for the first time.
	var err error
	db, err = sql.Open(dbDriver, dbName)
	if err != nil {
		log.Fatal(err)
	}
	// Ping the database to establish an actual first connection (to test only)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database: first connection successfully established")
}

// Create creates the database with the model provided in the package
func Create() {
	// Create train table
	_, err := db.Exec(createNotesSQL)
	if err != nil {
		log.Fatal(err)
	}
}
