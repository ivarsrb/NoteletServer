package database

import (
	"database/sql"
	"log"

	// To initialize sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Db database connection
var db *sql.DB

func init() {
	// Pprepares the database abstraction for later use
	// The first actual connection to the underlying datastore will
	// be established lazily, when it’s needed for the first time.
	const dbName = "./notelet.db"
	const dbDriver = "sqlite3"
	var err error
	db, err = sql.Open(dbDriver, dbName)
	if err != nil {
		log.Fatal(err)
	}
	// Ping the database to establish an actual connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
}

// Create creates the database with given model
func Create() {
	statement, driverError := db.Prepare(notesQuery)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}
}
