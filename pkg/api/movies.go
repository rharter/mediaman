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
func MovieList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return database.ListMovies()
}

// GET /api/movies/:id
// Shows details for a movie with :id
func MovieShow(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	m, err := database.GetMovie(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound{}
		} else {
			return nil, err
		}
	}

	return m, nil
}

// PUT /api/movies
// Creates a new movie
func MovieCreate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var m Movie
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return nil, err
	}

	if m.ID != 0 {
		return nil, badRequest{errors.New("Movie already has an ID")}
	}

	err = database.SaveMovie(&m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
