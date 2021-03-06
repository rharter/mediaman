package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

// badRequest is handled by setting the status code in the reply to StatusBadRequest
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

// ErrorHandler wraps the default http.HandleFunc to handl an
// error as the return value
type ErrorHandler func(w http.ResponseWriter, r *http.Request) error

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err == nil {
		return
	}
	switch err.(type) {
	case badRequest:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case notFound:
		http.Error(w, "File not found", http.StatusNotFound)
	default:
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}

func JsonHandler(f func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.Handler {
	return ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		d, err := f(w, r)
		if err != nil {
			return err
		}

		r.Header.Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(d)
		return err
	})
}

func AddHandlers(m *pat.PatternServeMux, base string) {

	// Movies
	m.Get(base+"/movies", JsonHandler(MovieList))
	m.Get(base+"/movies/:id", JsonHandler(MovieShow))
	m.Put(base+"/movies", JsonHandler(MovieCreate))
	m.Del(base+"/movies/:id", JsonHandler(MovieDelete))

	// Libraries
	m.Get(base+"/libraries", JsonHandler(LibraryList))
	m.Get(base+"/libraries/:id", JsonHandler(LibraryShow))
	m.Put(base+"/libraries", JsonHandler(LibraryCreate))
	m.Del(base+"/libraries/:id", JsonHandler(LibraryDelete))
	m.Get(base+"/libraries/:id/process", JsonHandler(LibraryProcess))

	// Series
	m.Get(base+"/series", JsonHandler(SeriesList))
	m.Get(base+"/series/:id", JsonHandler(SeriesShow))
	m.Put(base+"/series", JsonHandler(SeriesCreate))
	m.Del(base+"/series/:id", JsonHandler(SeriesDelete))

	// Series season
	m.Get(base+"/series/:id/seasons", JsonHandler(SeasonsList))

	// Series Episodes
	m.Get(base+"/series/:series_id/seasons/:season_num/episodes", JsonHandler(EpisodeList))

	// Episode
	m.Get(base+"/episodes/:id", JsonHandler(EpisodeShow))
	m.Put(base+"/episodes", JsonHandler(EpisodeCreate))
	m.Del(base+"/episodes/:id", JsonHandler(EpisodeDelete))
	m.Get(base+"/series/:id/episodes", JsonHandler(SeriesEpisodeList))
}
