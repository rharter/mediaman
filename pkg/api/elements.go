package api

import (
	"net/http"

	"github.com/rharter/mediaman/pkg/database"
)

// GET /elements/:id/children
// returns the children of a directory
func ElementList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	id, err := getIdFromRequest(r)
	if err != nil {
		return nil, err
	}

	return database.GetElementsForParent(id)
}

// GET /elements/:id/
// returns the details of an item
func ElementShow(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	id, err := getIdFromRequest(r)
	if err != nil {
		return nil, err
	}

	return database.GetElement(id)
}
