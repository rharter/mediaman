package handler

import (
	"net/http"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

// Display a movie list
func SeriesList(w http.ResponseWriter, r *http.Request) error {
	s, err := database.ListSeries()
	if err != nil {
		return err
	}

	data := struct {
		Series []*Series
	}{s}

	return RenderTemplate(w, r, "series_list.html", &data)
}
