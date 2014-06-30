package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Movie table in the database
const seriesTable = "series"

// SQL Query to retrieve a series by it's unique database key
const seriesFindIdStmt = `
SELECT id, language, title, overview, banner, poster, fanart, imdb_id, 
	series_id, created, updated
FROM series WHERE id = ?
`

const seriesByNameStmt = `
SELECT id, language, title, overview, banner, poster, fanart, imdb_id,
	series_id, created, updated
FROM series WHERE title = ?
`

// SQL Query to retrieve a series by it's path name
const seriesFindPathStmt = `
SELECT id, language, title, overview, banner, poster, fanart, imdb_id, 
	series_id, created, updated
FROM series WHERE filename = ?
`

// SQL Query to retrieve a all seriess
const seriesStmt = `
SELECT id, language, title, overview, banner, poster, fanart, imdb_id, 
	series_id, created, updated
FROM series
ORDER BY title ASC
`

// Returns the series with the given ID.
func GetSeries(id int64) (*Series, error) {
	series := Series{}
	err := meddler.QueryRow(db, &series, seriesFindIdStmt, id)
	return &series, err
}

func GetSeriesByTitle(t string) (*Series, error) {
	series := Series{}
	err := meddler.QueryRow(db, &series, seriesByNameStmt, t)
	return &series, err
}

// Saves a Series
func SaveSeries(series *Series) error {
	if series.Id == 0 {
		series.Created = time.Now().UTC()
	}
	series.Updated = time.Now().UTC()
	return meddler.Save(db, seriesTable, series)
}

// Deletes an existing Series.
func DeleteSeries(id int64) error {
	db.Exec("DELETE FROM series WHERE id = ?", id)
	return nil
}

// Lists all Series
func ListSeries() ([]*Series, error) {
	var series []*Series
	err := meddler.QueryAll(db, &series, seriesStmt)
	return series, err
}
