package config

import (
	"database/sql"
	"fmt"
	"os"
)

var CONFIGURATIONS map[string]string = map[string]string{
	"DATABASE_URL": "dbname=server sslmode=disable",
}

func InitDB() (*sql.DB, error) {
	if os.Getenv("DATABASE_URL") == "" {
		os.Setenv("DATABASE_URL", CONFIGURATIONS["DATABASE_URL"])
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Initialized read-write database connection pool: %v", os.Getenv("DATABASE_URL"))
	return db, nil
}
