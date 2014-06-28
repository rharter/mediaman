package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/rharter/mediaman/pkg/database"
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
