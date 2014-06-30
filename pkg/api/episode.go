package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

// GET /api/episodes
// Lists all episodes
func EpisodeList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	sidstr := getPathParam(r, "series_id")
	series_id, err := strconv.ParseInt(sidstr, 10, 64)
	if err != nil {
		return nil, err
	}

	snumstr := getPathParam(r, "season_num")
	season_num, err := strconv.ParseInt(snumstr, 10, 64)
	if err != nil {
		return nil, err
	}

	return database.ListEpisodesForSeriesSeason(series_id, season_num)
}

func SeriesEpisodeList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	return database.ListEpisodesForSeries(id)
}

// GET /api/episodes/:id
// Returns a single Episode identified by :id
func EpisodeShow(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	s, err := database.GetEpisode(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound{}
		} else {
			return nil, err
		}
	}

	return s, nil
}

// PUT /api/episodes
// Creates a new episodes
func EpisodeCreate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var s Episode
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	if s.Id != 0 {
		return nil, badRequest{errors.New("Episode already has an Id")}
	}

	err = database.SaveEpisode(&s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DELETE /api/episodes/:id
// Delete an episode identified by :id
func EpisodeDelete(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	s, err := database.GetEpisode(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound{}
		} else {
			return nil, err
		}
	}

	err = database.DeleteEpisode(s.Id)
	if err != nil {
		return nil, err
	}

	return s, nil
}
