package tmdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type Movie struct {
	Adult        bool   `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	//BelongsToCollection something `json:"belongs_to_collection"`
	Budget              int64     `json:"budget"`
	Genres              []Genre   `json:"genres"`
	Homepage            string    `json:"homepage"`
	ID                  int64     `json:"id"`
	ImdbID              string    `json:"imdb_id"`
	OriginalTitle       string    `json:"original_title"`
	Overview            string    `json:"overview"`
	Popularity          float64   `json:"popularity"`
	PosterPath          string    `json:"poster_path"`
	ProductionCompanies []Company `json:"production_companies"`
	ProductionCountries []Country `json:"production_countries"`
	// 	ReleaseDate         time.Time `json:"release_date"`
	Revenue         int64      `json:"revenue"`
	Runtime         int64      `json:"runtime"`
	SpokenLanguages []Language `json:"spoken_languages"`
	Tagline         string     `json:"tagline"`
	Title           string     `json:"title"`
	VoteAverage     float64    `json:"vote_average"`
	VoteCount       int64      `json:"vote_count"`
}

type Language struct {
	Name    string `json:"name"`
	IsoCode string `json:"iso_639_1"`
}

type Company struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Country struct {
	Name    string `json:"name"`
	IsoCode string `json:"iso_3166_1"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Gets a Movie by ID
func (t *TMDB) GetMovie(id int64) (*Movie, error) {
	resp, err := t.performGet(fmt.Sprintf("movie/%d", id), nil)
	if err != nil {
		return nil, err
	}

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (t *TMDB) SearchMovies(query string) (movies []Movie, err error) {
	args := url.Values{}
	args.Set("query", query)

	resp, err := t.performGet("search/movie", args)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Failed to parse response: %s", err)
		return nil, err
	}

	err = json.Unmarshal(result.ResultsRaw, &movies)

	return
}
