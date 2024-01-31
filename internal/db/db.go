package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=medicloud_postgres sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	log.Println("Attempting to connect to the database..." + connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error connecting to the database: %v\n", err)
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return db, nil
}
