package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

// GET /libraries
// returns a list of libraries
func LibraryList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return database.ListLibraries()
}

// PUT /libraries
// Creates a new library and returns it
func LibraryCreate(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var l Library
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		return nil, err
	}

	if l.ID != 0 {
		return nil, badRequest{errors.New("Library already has an ID")}
	}

	err = database.SaveLibrary(&l)
	if err != nil {
		return nil, err
	}

	return l, nil
}
