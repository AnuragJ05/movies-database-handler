package main

import (
	"log"
	"net/http"

	config "github.com/AnuragJ05/database-handler/config"
	util "github.com/AnuragJ05/database-handler/util"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	//connect to database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		log.Fatal(err)
	}

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/users", util.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", util.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", util.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", util.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", util.DeleteUser(db)).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
