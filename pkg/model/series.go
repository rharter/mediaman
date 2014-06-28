package model

import (
	"time"

	"github.com/theraphim/tvdb"
)

type Series struct {
	Id       int64  `meddler:"id,pk"     json:"id"`
	Language string `meddler:"language"  json:"language"`
	Title    string `meddler:"title"     json:"title"`
	Overview string `meddler:"overview"  json:"overview"`
	Banner   string `meddler:"banner"    json:"banner"`
	ImdbId   string `meddler:"imdb_id"   json:"imdb_id"`
	SeriesId int64  `meddler:"series_id" json:"series_id"`

	Created time.Time `meddler:"created,utctime" json:"created_at"`
	Updated time.Time `meddler:"updated,utctime" json:"updated_at"`
}

func NewSeriesFromTvdb(source tvdb.Series) *Series {
	return &Series{
		Id:       int64(source.Id),
		Language: source.Language,
		Title:    source.SeriesName,
		Overview: source.Overview,
		Banner:   source.Banner,
		ImdbId:   source.ImdbId,
		SeriesId: int64(source.SeriesId),
	}
}
