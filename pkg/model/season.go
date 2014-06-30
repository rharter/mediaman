package model

import (
	"errors"
	"time"
)

type Season struct {
	Id           int64     `meddler:"id,pk" 		  json:"id"`
	SeasonNumber uint64    `meddler:"season_number"   json:"season_number"`
	SeriesId     int64     `meddler:"series_id" 	  json:"series_id"`
	Created      time.Time `meddler:"created,utctime" json:"created"`
	Updated      time.Time `meddler:"updated,utctime" json:"updated"`

	// Transient
	Episodes []*Episode `meddler:"-" json:"episodes"`
}

func NewSeason(seriesId int64, number uint64) (*Season, error) {
	if seriesId == 0 || number == 0 {
		return nil, errors.New("seriesId and number are required")
	}
	return &Season{
		SeriesId:     seriesId,
		SeasonNumber: number,
	}, nil
}
