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

	model "movies-database-handler/model"

	_ "github.com/lib/pq"
)

// saveToFile saves the given data to a file at the specified file path.
//
// Parameters:
// - data: a byte slice containing the data to be written to the file.
// - filePath: a string representing the path of the file to be created or overwritten.
//
// Returns:
// - error: an error if the file creation or writing operation fails.
func saveToFile(data []byte, filePath string) error {

	// Create the directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Write the data to the file with 0644 permissions
	return ioutil.WriteFile(filePath, data, 0644)
}

// GetMovies returns an HTTP handler function that retrieves all movies from the
// database and encodes them as JSON in the response body. It takes a pointer to a
// sql.DB object as a parameter. The function does not return anything.
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models)

	}
}

// CreateMovie handles the creation of a new movie.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// The function sets the content type of the response to "application/json".
// It then decodes the JSON data from the request body into a Movie struct.
// The function generates a new movie ID using the rand.Intn function and converts it to a string.
// It sets the Timestamp field of the movie struct to the current time in RFC3339 format.
// The function encodes the movie struct into JSON and writes it to the response.
// It creates a file path using the movie ID and saves the movie struct as JSON in that file.
// If there is an error during the process, it returns an appropriate HTTP error response.
func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie

	// decoding the JSON data from the request body into a Movie struct
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Print(movie)

	// generating a new movie ID using the rand.Intn function and converting it to a string
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movie.Timestamp = time.Now().Format(time.RFC3339)

	json.NewEncoder(w).Encode(movie)

	filePath := fmt.Sprintf("/tmp/astra/%s.json", movie.ID)

	body, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// creating a file path using the movie ID and saving the movie struct as JSON in that file
	if err := saveToFile(body, filePath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
