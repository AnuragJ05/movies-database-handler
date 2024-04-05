package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	model "github.com/AnuragJ05/database-handler/model"
	_ "github.com/lib/pq"
)

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

			// Open JSON file
			file, err := os.Open(latestFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			// Read JSON data from file
			jsonData, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}

			// Unmarshal JSON data into a slice of Movie structs
			var movie model.Movie
			if err := json.Unmarshal(jsonData, &movie); err != nil {
				log.Fatal(err)
			}
			_, err = db.Exec("INSERT INTO movies (id, isbn, title, director, timestamp) VALUES ($1, $2, $3, $4, $5)",
				movie.ID, movie.Isbn, movie.Title, movie.Director, movie.Timestamp)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Inserted:", movie.Title)

		}

	}

}

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
		fmt.Println("Error:", err)
		return
	}

	ch <- latestFileName
}
