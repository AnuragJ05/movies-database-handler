package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/AnuragJ05/database-handler/config"
	"github.com/AnuragJ05/database-handler/util"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// main is the entry point of the Go program.
//
// It connects to the database, creates a table if it doesn't exist, and sets up
// the router to handle different HTTP methods for the "/users" endpoint.
// Finally, it starts the server and listens for incoming requests on port 8000.
func main() {

	// port := os.Getenv("PORT")

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

	//create router
	r := mux.NewRouter() // function from gorila
	r.HandleFunc("/movies", util.GetMovies(db)).Methods("GET")
	r.HandleFunc("/movies", util.CreateMovie).Methods("POST")

	fmt.Println("Starting serve at port /5000")
	go http.ListenAndServe(":5000", r)

	var wg sync.WaitGroup
	wg.Add(1)
	go util.UpdateDBFromFile(&wg, db)

	wg.Wait()

	fmt.Println("END")
}
