package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
	"github.com/rharter/mediaman/pkg/transcoder"
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

	return RenderTemplate(w, r, "movies_list.html", &data)
}

// /movies/:id
// Shows movie details for the movie specified by :id
func MovieShow(w http.ResponseWriter, r *http.Request) error {
	id, err := movieIdFromRequest(r)
	if err != nil {
		return err
	}

	movie, err := database.GetMovie(id)
	if err != nil {
		return err
	}

	return RenderTemplate(w, r, "movies_show.html", movie)
}

func MovieTranscode(w http.ResponseWriter, r *http.Request) {
	id, err := movieIdFromRequest(r)
	if err != nil {
		return
	}

	movie, err := database.GetMovie(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	log.Printf("Starting transcoding session for file: %s", movie.Filename)
	trans := transcoder.NewTranscoderSession(movie.ID, "/tmp/videos", movie.Filename)

	if err = trans.Open(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rel, err := filepath.Rel("/tmp/videos", trans.OutputFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	p := filepath.Join("/videos/", rel)

	log.Printf("Serving output file: %s", p)
	http.Redirect(w, r, p, http.StatusTemporaryRedirect)
}

// Streams the actual video file directly.
func MoviePlay(w http.ResponseWriter, r *http.Request) {
	id, err := movieIdFromRequest(r)
	if err != nil {
		return
	}

	movie, err := database.GetMovie(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, movie.Filename)
}

func movieIdFromRequest(r *http.Request) (id int64, err error) {
	idstr := r.URL.Query().Get(":id")
	return strconv.ParseInt(idstr, 10, 64)
}
