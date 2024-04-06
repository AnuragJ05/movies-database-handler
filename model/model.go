package model

// Movie represents a movie in the database table
type Movie struct {
	ID        string `json:"id"`
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Director  string `json:"director"`
	Timestamp string `json:"timestamp"`
}
