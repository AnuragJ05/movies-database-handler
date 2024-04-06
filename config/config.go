package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// InitDB initializes a database connection pool for read-write operations.
//
// It retrieves the database connection details from environment variables
// (DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD) and constructs the
// database connection string. It then opens a connection to the database
// using the "postgres" driver and returns the database handle and any
// error encountered during the initialization process.
//
// Returns:
// - *sql.DB: The database handle for read-write operations.
// - error: An error if the database connection fails.
func InitDB() (*sql.DB, error) {

	// Get database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	// Construct the database connection string
	connStrMain := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStrMain)

	if err != nil {
		return nil, err
	}

	log.Printf("Initialized read-write database connection pool: %v", os.Getenv("DATABASE_URL"))
	return db, nil
}
