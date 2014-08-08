package processor

import (
	"log"
	"path/filepath"

	"github.com/rharter/go-guessit"
	"github.com/rharter/go-tmdb"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

const (

	// The TMDB API Key
	TMDB_API_KEY = "082009c4d5bcbc92baea080023ae8b3d"
)

var con *tmdb.TMDB

func init() {
	con = tmdb.Open(TMDB_API_KEY)
}

type FetchMovieMetadataTask struct {
	Element *Element
}

// FetchMetadataTask interface
func (t *FetchMovieMetadataTask) Fetch() error {

	// See if we already have a element for this file
	element, err := database.GetElementByFile(t.Element.File)
	if err != nil {
		element = t.Element
	}

	var query string
	if element.Title != "" {
		query = element.Title
	} else {
		// Try to get the file name info from GuessIt
		meta, err := guessit.Guess(element.File)
		if err != nil {
			log.Printf("Failed to guess element info: %+v", err)

			// Fallback to simple filename guessing
			filename := filepath.Base(element.File)
			extension := filepath.Ext(element.File)
			query = filename[:len(filename)-len(extension)]
		} else {
			query = meta.Title
		}
	}
	log.Printf("Fetching metadata for %s", query)

	movies, err := con.SearchMovies(query)
	if err != nil {
		log.Printf("Error querying TheMovieDB for %s: %+v", query, err)
		return err
	}

	if len(movies) < 1 {
		log.Printf("No movies found for name: %s", query)
		return nil
	}

	// For now, we'll assume the first match is the winner
	match := movies[0]
	element.Title = match.Title
	element.Background = match.BackdropPath
	element.Poster = match.PosterPath
	element.Description = match.Overview
	err = database.SaveElement(element)
	if err != nil {
		return err
	}

	return nil
}
