package tmdb

import (
	"log"
	"net/http"
	"net/url"
)

const (
	// The base API path
	API_BASE = "http://api.themoviedb.org/3/"
)

type TMDB struct {
	// Location to use for API access
	Location string

	// Api key
	ApiKey string
}

func Open(apiKey string) *TMDB {
	return &TMDB{
		Location: API_BASE,
		ApiKey:   apiKey,
	}
}

func (t *TMDB) performGet(path string, args url.Values) (*http.Response, error) {
	if args == nil {
		args = url.Values{}
	}
	args.Set("api_key", t.ApiKey)

	query, err := url.ParseRequestURI(t.Location)
	if err != nil {
		return nil, err
	}

	if query, err = query.Parse(path); err != nil {
		return nil, err
	}

	query.RawQuery = args.Encode()
	log.Printf("Executing request: %s", query.String())
	return http.Get(query.String())
}
