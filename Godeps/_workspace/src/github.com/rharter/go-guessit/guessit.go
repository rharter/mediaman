package guessit

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	BASE_URL = "http://guessit.io/"
)

type GuessResult struct {
	AudioChannels string `json:"audioChannels"`
	AudioCodec    string `json:"audioCodec"`
	Container     string `json:"container"`
	EpisodeNumber int64  `json:"episodeNumber"`
	Format        string `json:"format"`
	MimeType      string `json:"mimetype"`
	ReleaseGroup  string `json:"releaseGroup"`
	ScreenSize    string `json:"screenSize"`
	Season        int64  `json:"season"`
	Series        string `json:"series"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	VideoCodec    string `json:"videoCodec"`
	Year          int    `json:"year"`
}

func Guess(filename string) (*GuessResult, error) {
	args := url.Values{}
	args.Set("filename", filename)

	query, err := url.ParseRequestURI(BASE_URL)
	if err != nil {
		return nil, err
	}

	query, err = query.Parse("/guess")
	if err != nil {
		return nil, err
	}

	query.RawQuery = args.Encode()

	r, err := http.Get(query.String())
	if err != nil {
		return nil, err
	}

	var guess GuessResult
	err = json.NewDecoder(r.Body).Decode(&guess)
	if err != nil {
		return nil, err
	}

	return &guess, err
}
