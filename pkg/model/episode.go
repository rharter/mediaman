package model

import (
	"errors"
	"time"

	"github.com/rharter/tvdb"
)

type Episode struct {
	Id             int64     `meddler:"id,pk"           json:"id"`
	Title          string    `meddler:"title"           json:"title"`
	Overview       string    `meddler:"overview"        json:"overview"`
	Director       string    `meddler:"director"        json:"director"`
	Writer         string    `meddler:"writer"          json:"writer"`
	GuestStars     string    `meddler:"guest_stars"     json:"guest_stars"`
	SeasonId       int64     `meddler:"season_id"       json:"season_id"`
	EpisodeNumber  uint64    `meddler:"episode_number"  json:"episode_number"`
	SeasonNumber   uint64    `meddler:"season_number"   json:"season_number"`
	AbsoluteNumber string    `meddler:"absolute_number" json:"absolute_number"`
	Language       string    `meddler:"language"        json:"language"`
	Rating         string    `meddler:"rating"          json:"rating"`
	SeriesId       int64     `meddler:"series_id"       json:"series_id"`
	TvdbSeriesId   int64     `meddler:"tvdb_series_id"  json:"tvdb_series_id"`
	ImdbId         string    `meddler:"imdb_id"         json:"imdb_id"`
	Poster         string    `meddler:"poster"          json:"poster"`
	Filename       string    `meddler:"filename"        json:"filename"`
	Created        time.Time `meddler:"created,utctime" json:"created"`
	Updated        time.Time `meddler:"updated,utctime" json:"updated"`
}

func NewEpisode(filename string) (*Episode, error) {
	if filename == "" {
		return nil, errors.New("empty filename")
	}
	return &Episode{Filename: filename}, nil
}

func EpisodeFromTvdbEpisode(e *tvdb.Episode) *Episode {
	return &Episode{
		Title:          e.EpisodeName,
		Overview:       e.Overview,
		Director:       e.Director,
		Writer:         e.Writer,
		GuestStars:     e.GuestStars,
		SeasonId:       int64(e.SeasonId),
		EpisodeNumber:  uint64(e.EpisodeNumber),
		SeasonNumber:   uint64(e.SeasonNumber),
		AbsoluteNumber: e.AbsoluteNumber,
		TvdbSeriesId:   int64(e.SeriesId),
		Language:       e.Language,
		Rating:         e.Rating,
		ImdbId:         e.ImdbId,
	}
}
