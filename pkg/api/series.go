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

// GET /api/movies
// Lists all movies
func SeriesList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return database.ListSeries()
}

// GET /api/series/:id
// Returns a single Series identified by :id
func SeriesShow(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	s, err := database.GetSeries(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound{}
		} else {
			return nil, err
		}
	}

	return s, nil
}

// PUT /api/series
// Creates a new series
func SeriesCreate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var s Series
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	if s.Id != 0 {
		return nil, badRequest{errors.New("Series already has an Id")}
	}

	err = database.SaveSeries(&s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DELETE /api/series/:id
// Delete a series identified by :id
func SeriesDelete(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	s, err := database.GetSeries(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound{}
		} else {
			return nil, err
		}
	}

	err = database.DeleteSeries(s.Id)
	if err != nil {
		return nil, err
	}

	return s, nil
}
