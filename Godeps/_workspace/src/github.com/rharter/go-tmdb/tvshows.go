package tmdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type TvShow struct {
	Id           int64   `json:"id"`
	BackdropPath string  `json:"backdrop_path"`
	OriginalName string  `json:"original_name"`
	FirstAirDate string  `json:"first_air_date"`
	PosterPath   string  `json:"poster_path"`
	Popularity   float64 `json:"popularity"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
}

type Cast struct {
	Id          int64  `json:"id"`
	Character   string `json:"character"`
	CreditId    string `json:"credit_id"`
	Name        string `json:"name"`
	ProfilePath string `json:"profile_path"`
	Order       int    `json:"order"`
}

type Crew struct {
	Id          int64  `json:"id"`
	Department  string `json:"department"`
	Name        string `json:"name"`
	Job         string `json:"job"`
	ProfilePath string `json:"profile_path"`
}

type Credit struct {
	Id   int64  `json:"id"`
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

// Get the primary information about a TV series by id.
func (t *TMDB) GetTvShowById(id int64) (*TvShow, error) {
	resp, err := t.performGet(fmt.Sprintf("tv/%d", id), nil)
	if err != nil {
		return nil, err
	}

	var show TvShow
	err = json.NewDecoder(resp.Body).Decode(&show)
	if err != nil {
		return nil, err
	}

	return &show, nil
}

// Get the cast & crew information about a TV series.
func (t *TMDB) GetTvShowCredits(id int64) (*Credit, error) {
	resp, err := t.performGet(fmt.Sprintf("tv/%d/credits", id), nil)
	if err != nil {
		return nil, err
	}

	var c Credit
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (t *TMDB) SearchTvShows(query string) (shows []TvShow, err error) {
	args := url.Values{}
	args.Set("query", query)

	resp, err := t.performGet("search/tv", args)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Failed to parse response: %s", err)
		return nil, err
	}

	err = json.Unmarshal(result.ResultsRaw, &shows)

	return
}
