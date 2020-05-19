package database

import (
	"log"
)

// Initialize Create the database with given model
func Initialize() {
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
