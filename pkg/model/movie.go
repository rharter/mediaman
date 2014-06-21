package model

import (
	"errors"
	"time"
)

type Movie struct {
	ID           int64     `meddler:"id,pk"            json:"id"`
	Title        string    `meddler:"title"            json:"title"`
	BackdropPath string    `meddler:"backdrop"         json:"backdrop"`
	PosterPath   string    `meddler:"poster"           json:"poster"`
	ReleaseDate  time.Time `meddler:"release_date"     json:"release_date"`
	Adult        bool      `meddler:"adult"            json:"adult"`
	Genres       string    `meddler:"genres"           json:"genres"`
	Homepage     string    `meddler:"homepage"         json:"homepage"`
	ImdbID       string    `meddler:"imdb_id"          json:"imdb_id"`
	Overview     string    `meddler:"overview"         json:"overview"`
	Runtime      int64     `meddler:"runtime"          json:"runtime"`
	Tagline      string    `meddler:"tagline"          json:"tagline"`
	UserRating   float64   `meddler:"rating"           json:"rating"`
	Created      time.Time `meddler:"created,utctime"  json:"created"`
	Updated      time.Time `meddler:"updated,utctime"  json:"updated"`

	Filename string `meddler:"filename"       json:"filename"`
}

func NewMovie(filename string) (*Movie, error) {
	if filename == "" {
		return nil, errors.New("empty filename")
	}
	return &Movie{Filename: filename}, nil
}
