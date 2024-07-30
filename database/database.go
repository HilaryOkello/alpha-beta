package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./orders.db")
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	// Create orders table if it does not exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		manufacturer_id TEXT,
		health_facility_id TEXT,
		vaccine_details TEXT,
		transaction_type TEXT,
		status TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("could not create table: %v", err)
	}
}
