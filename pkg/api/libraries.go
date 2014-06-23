package api

import (
	"net/http"

	"github.com/rharter/mediaman/pkg/database"
)

func LibraryList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return database.ListLibraries()
}
