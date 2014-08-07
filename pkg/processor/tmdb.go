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
	Video *Video
}

// FetchMetadataTask interface
func (t *FetchMovieMetadataTask) Fetch() error {

	// See if we already have a video for this file
	video, _ := database.GetVideoByFile(t.Video.File)
	if video.Title != "" {
		return nil
	}

	var query string
	if t.Video.Title != "" {
		query = t.Video.Title
	} else {
		// Try to get the file name info from GuessIt
		meta, err := guessit.Guess(t.Video.File)
		if err != nil {
			log.Printf("Failed to guess video info: %+v", err)

			// Fallback to simple filename guessing
			filename := filepath.Base(t.Video.File)
			extension := filepath.Ext(t.Video.File)
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
	t.Video.Title = match.Title
	t.Video.Background = match.BackdropPath
	t.Video.Poster = match.PosterPath
	t.Video.Description = match.Overview
	err = database.SaveVideo(t.Video)
	if err != nil {
		return err
	}

	return nil
}
