package config

import (
	"database/sql"
	"fmt"
	"os"
)

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

	// connStrMain := "user=postgres password=postgres dbname=postgres sslmode=disable port=5432 hostdb=postgres"
	db, err := sql.Open("postgres", connStrMain)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Initialized read-write database connection pool: %v", os.Getenv("DATABASE_URL"))
	return db, nil
}
