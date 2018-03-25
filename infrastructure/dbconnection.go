package infrastructure

import (
	"database/sql"
	"log"

 	_ "github.com/mattn/go-sqlite3"
 )

var databaseFile string = "./database.db";

func GetDbConnection () *sql.DB {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}