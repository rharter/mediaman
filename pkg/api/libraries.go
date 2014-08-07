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

// GET /libraries
// returns a list of libraries
func LibraryList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return database.ListLibraries()
}

// GET /api/libraries/:id
// Shows details for a library with :id
func LibraryShow(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	l, err := database.GetLibrary(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound{}
		} else {
			return nil, err
		}
	}

	return l, nil
}

// PUT /libraries
// Creates a new library and returns it
func LibraryCreate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var l Library
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		return nil, err
	}

	if l.Id != 0 {
		return nil, badRequest{errors.New("Library already has an Id")}
	}

	err = database.SaveLibrary(&l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

// DELETE /libraries/:id
// Deletes a library identified by :id
func LibraryDelete(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, err
	}

	l, err := database.GetLibrary(id)
	if err != nil {
		return nil, err
	}

	err = database.DeleteLibrary(id)
	if err != nil {
		return nil, err
	}

	return l, nil
}
