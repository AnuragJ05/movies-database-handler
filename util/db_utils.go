package util

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	model "movies-database-handler/model"

	_ "github.com/lib/pq"
)

// UpdateDBFromFile updates the database from a file.
//
// It takes a pointer to a sync.WaitGroup and a pointer to a sql.DB as parameters.
// The function continuously monitors the "/tmp/astra/" directory for the latest file created.
// When a new file is detected, it reads the JSON data from the file and inserts the data into the "movies" table in the database.
// The function logs the latest file created and the movie added to the database.
// The function runs indefinitely until it is stopped.
func UpdateDBFromFile(wgGlobal *sync.WaitGroup, db *sql.DB) {

	defer wgGlobal.Done()
	lastFile := ""
	for {
		dir := "/tmp/astra/"
		var wgLocal sync.WaitGroup
		ch := make(chan string)

		wgLocal.Add(1)
		go findLatestFile(dir, &wgLocal, ch)

		go func() {
			wgLocal.Wait()
			close(ch)
		}()

		latestFile := <-ch
		if lastFile != latestFile {
			lastFile = latestFile

			log.Println("Latest file created:", latestFile)

			// Open JSON file
			file, err := os.Open(latestFile)
			if err != nil {
				log.Fatal(err)
			}

			// Read JSON data from file
			jsonData, err := io.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}

			// Unmarshal JSON data into Movie struct
			var movie model.Movie
			if err := json.Unmarshal(jsonData, &movie); err != nil {
				log.Fatal(err)
			}
			_, err = db.Exec("INSERT INTO movies (id, isbn, title, director, timestamp) VALUES ($1, $2, $3, $4, $5)",
				movie.ID, movie.Isbn, movie.Title, movie.Director, movie.Timestamp)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			log.Println("Movie added to database:", movie.Title)

		}

	}

}

// findLatestFile finds the latest modified file in the given directory and sends its path to the provided channel.
//
// Parameters:
// - dir: the directory to search for the latest file.
// - wg: a pointer to a sync.WaitGroup used for synchronization.
// - ch: a channel of type string to send the path of the latest file.
func findLatestFile(dir string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	var latestModTime time.Time
	var latestFileName string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().After(latestModTime) {
			latestModTime = info.ModTime()
			latestFileName = path
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	ch <- latestFileName
}
