package tmdb

import (
	"encoding/json"
	"fmt"
)

type Season struct {
	Id           int64      `json:"id"`
	Name         string     `json:"name"`
	Overview     string     `json:"overview"`
	PosterPath   string     `json:"poster_path"`
	SeasonNumber int64      `json:"season_number"`
	AirDate      string     `json:"air_date"`
	Episodes     []*Episode `json:"episodes"`
}

// Get the primary information about a TV season by its
// season number.
func (t *TMDB) GetSeasonByNumber(i int64, n int64) (*Season, error) {
	resp, err := t.performGet(fmt.Sprintf("tv/%d/season/%d", i, n), nil)
	if err != nil {
		return nil, err
	}

	var s Season
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
