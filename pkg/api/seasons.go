package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rharter/mediaman/pkg/database"
)

// GET /series/:id/seasons
// Returns a list of seasons for a given series
func SeasonsList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	idstr := getPathParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unparsable id: %s", idstr))
	}

	return database.ListSeasonsForSeries(id)
}
