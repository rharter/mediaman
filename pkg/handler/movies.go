package handler

import (
	"net/http"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

// Display a movie list
func MovieList(w http.ResponseWriter, r *http.Request) error {
	movies, err := database.ListMovies()
	if err != nil {
		return err
	}

	data := struct {
		Movies []*Movie `json:"movies"`
	}{movies}

	return RenderTemplate(w, r, "list_movies.html", &data)
}