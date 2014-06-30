package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Movie table in the database
const (
	// The name of the table
	seasonTable = "seasons"

	// SQL Query to retrieve a season by it's unique database key
	seasonFindByIdStmt = `
		SELECT id, season_number, series_id, created, updated 
		FROM seasons 
		WHERE id = ?
	`
	// SQL Query to retrieve all seasons for a series
	seasonFindBySeriesStmt = `
		SELECT id, season_number, series_id, created, updated 
		FROM seasons 
		WHERE series_id = ?
		ORDER BY season_number ASC
	`

	// SQL Query to retrieve a specific season by number for a specific series
	seasonFindByNumberForSeriesStmt = `
		SELECT id, season_number, series_id, created, updated 
		FROM seasons 
		WHERE season_number = ? AND series_id = ?
		ORDER BY season_number ASC
	`

	// SQL Query to retrieve a all seriess
	seasonsStmt = `
		SELECT id, season_number, series_id, created, updated
		FROM seasons
		ORDER BY season_number ASC
	`
)

// Returns all season for a series
func ListSeasonsForSeries(id int64) ([]*Season, error) {
	var seasons []*Season
	err := meddler.QueryAll(db, &seasons, seasonFindBySeriesStmt, id)
	return seasons, err
}

// Returns a specific season from a series
func GetSeasonForSeries(seriesId int64, num uint64) (*Season, error) {
	var season Season
	err := meddler.QueryRow(db, &season, seasonFindByNumberForSeriesStmt, num, seriesId)
	return &season, err
}

// Saves a season
func SaveSeason(s *Season) error {
	if s.Id == 0 {
		s.Created = time.Now().UTC()
	}
	s.Updated = time.Now().UTC()
	return meddler.Save(db, seasonTable, s)
}

// Deletes a season
func DeleteSeason(id int64) error {
	_, err := db.Exec("DELETE FROM seasons WHERE id = ?", id)
	return err
}
