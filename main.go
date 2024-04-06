package main

import (
	"log"
	"movies-database-handler/config"
	"movies-database-handler/util"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// main is the entry point of the Go program.
//
// It initializes the database connection, creates the necessary table if it doesn't exist,
// creates a router for handling HTTP requests, starts the server on port 5000,
// and updates the database from a file in the background.
//
// The function does not take any parameters and does not return any values.
func main() {

	// connect to database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS movies (id SERIAL PRIMARY KEY, isbn TEXT, title TEXT, director TEXT, timestamp TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	//create router for handling HTTP requests
	r := mux.NewRouter()                                       // NewRouter returns a new Gorilla Mux router
	r.HandleFunc("/movies", util.GetMovies(db)).Methods("GET") // GetMovies is a handler function that returns a list of movies.
	r.HandleFunc("/movies", util.CreateMovie).Methods("POST")  // CreateMovie is a handler function that creates a new movie.

	log.Println("Server started on port 5000")
	go http.ListenAndServe(":5000", r)

	var wg sync.WaitGroup
	wg.Add(1)
	go util.UpdateDBFromFile(&wg, db)
	wg.Wait()

	log.Panicln("Server stopped on port 5000")
}
