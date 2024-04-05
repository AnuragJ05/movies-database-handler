package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	model "github.com/AnuragJ05/database-handler/model"
	_ "github.com/lib/pq"
)

func saveToFile(data []byte, filePath string) error {

	// os.Create(filePath)

	// Create the directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Write the data to the file
	return ioutil.WriteFile(filePath, data, 0644)
}

// GetMovies is a handler function that returns a list of movies.
func GetMovies(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT * FROM movies")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		models := []model.Movie{}
		for rows.Next() {
			var m model.Movie
			if err := rows.Scan(&m.ID, &m.Isbn, &m.Title, &m.Director, &m.Timestamp); err != nil {
				log.Fatal(err)
			}
			models = append(models, m)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json") //Setting it bcz we want to work with json
		json.NewEncoder(w).Encode(models)

	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie

	// sending the data in the body as it is large and url does not have any parts of it
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Print(movie)

	// getting new movie id from random function which is integer
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // converting to string
	movie.Timestamp = time.Now().Format(time.RFC3339)

	json.NewEncoder(w).Encode(movie)

	filePath := fmt.Sprintf("/tmp/astra/%s.json", movie.ID)

	// Convert movie struct to JSON byte array
	body, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := saveToFile(body, filePath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
