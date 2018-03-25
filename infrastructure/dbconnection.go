// Package infrastructure package that provides an infrastructure to application
// Eg: databases, feature toogles
package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var databaseFile string = "./database.db"

// GetDbConnection just a function that create a db connection
func GetDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
