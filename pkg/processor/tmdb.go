package processor

import (
	"path/filepath"
	"log"

	"github.com/rharter/go-tmdb"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/rharter/mediaman/pkg/database"
)

const (

	// The TMDB API Key
	TMDB_API_KEY = "082009c4d5bcbc92baea080023ae8b3d"

)

var con *tmdb.TMDB

func init() {
	con = tmdb.Open(TMDB_API_KEY)
}

func FetchMetadataForMovie(movie *Movie) (err error) {
	var query string
	if movie.Title != "" {
		query = movie.Title
	} else {
		filename := filepath.Base(movie.Filename)
		extension := filepath.Ext(movie.Filename)
		query = filename[:len(filename)-len(extension)]
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
	movie.Title = match.Title
	movie.BackdropPath = match.BackdropPath
	movie.PosterPath = match.PosterPath
	movie.Adult = match.Adult
	movie.Homepage = match.Homepage
	movie.ImdbID = match.ImdbID
	movie.Overview = match.Overview
	movie.Runtime = match.Runtime
	movie.Tagline = match.Tagline
	movie.UserRating = match.VoteAverage
	err = database.SaveMovie(movie)
	if err != nil {
		return err
	}

	return nil
}