package config

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDB() (*sql.DB, error) {

	connStrMain := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStrMain)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Initialized read-write database connection pool: %v", os.Getenv("DATABASE_URL"))
	return db, nil
}
