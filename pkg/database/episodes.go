package database

import (
	"database/sql"
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Movie table in the database
const episodesTable = "episodes"

// SQL Query to retrieve a episodes by it's unique database key
const episodesFindIdStmt = `
SELECT id, title, overview, director, writer, guest_stars, 
	season_id, episode_number, season_number, absolute_number, 
	language, rating, series_id, imdb_id, filename, poster, created, updated
FROM episodes WHERE id = ?
`

// SQL Query to retrieve a episodes by it's SeriesId
const episodesFindSeriesIdStmt = `
SELECT id, title, overview, director, writer, guest_stars, 
	season_id, episode_number, season_number, absolute_number, 
	language, rating, series_id, imdb_id, filename, poster, created, updated
FROM episodes WHERE series_id = ?
`

// SQL Query to retrieve a episodes by it's SeriesId and season number
const episodesFindSeriesIdSeasonStmt = `
SELECT id, title, overview, director, writer, guest_stars, 
	season_id, episode_number, season_number, absolute_number, 
	language, rating, series_id, imdb_id, filename, poster, created, updated
FROM episodes WHERE series_id = ? AND season_number = ?
`

// SQL Query to retrieve a episodes by it's SeriesId and season number
const episodesFindFilenameStmt = `
SELECT id, title, overview, director, writer, guest_stars, 
	season_id, episode_number, season_number, absolute_number, 
	language, rating, series_id, imdb_id, filename, poster, created, updated
FROM episodes WHERE filename = ?
`

// SQL Query to retrieve a all episodess
const episodesStmt = `
SELECT id, title, overview, director, writer, guest_stars, 
	season_id, episode_number, season_number, absolute_number, 
	language, rating, series_id, imdb_id, filename, poster, created, updated
FROM episodes
ORDER BY title ASC
`

// Returns the Episode with the given ID.
func GetEpisode(id int64) (*Episode, error) {
	episode := Episode{}
	err := meddler.QueryRow(db, &episode, episodesFindIdStmt, id)
	return &episode, err
}

func GetEpisodeByFilename(n string) (*Episode, error) {
	episode := Episode{}
	err := meddler.QueryRow(db, &episode, episodesFindFilenameStmt, n)
	return &episode, err
}

// Saves an Episode
func SaveEpisode(e *Episode) error {
	// Save the season if we don't have one
	s, err := GetSeasonForSeries(e.SeriesId, e.SeasonNumber)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		s = &Season{SeriesId: e.SeriesId, SeasonNumber: e.SeasonNumber}
		err := SaveSeason(s)
		if err != nil {
			return err
		}

		e.SeasonId = s.Id
	}

	if e.Id == 0 {
		e.Created = time.Now().UTC()
	}
	e.Updated = time.Now().UTC()
	return meddler.Save(db, episodesTable, e)
}

// Deletes an existing Episode.
func DeleteEpisode(id int64) error {
	db.Exec("DELETE FROM episodes WHERE id = ?", id)
	return nil
}

// Lists all Episodes
func ListEpisodes() ([]*Episode, error) {
	var episodes []*Episode
	err := meddler.QueryAll(db, &episodes, episodesStmt)
	return episodes, err
}

func ListEpisodesForSeries(id int64) ([]*Episode, error) {
	var episodes []*Episode
	err := meddler.QueryAll(db, &episodes, episodesFindSeriesIdStmt, id)
	return episodes, err
}

func ListEpisodesForSeriesSeason(seriesId, seasonNum int64) ([]*Episode, error) {
	var episodes []*Episode
	err := meddler.QueryAll(db, &episodes, episodesFindSeriesIdSeasonStmt, seriesId, seasonNum)
	return episodes, err
}
