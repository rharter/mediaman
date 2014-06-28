package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Movie table in the database
const movieTable = "movies"

// SQL Query to retrieve a movie by it's unique database key
const movieFindIdStmt = `
SELECT id, title, backdrop, poster, release_date, adult, genres, 
	homepage, imdb_id, overview, runtime, tagline, rating, filename
FROM movies WHERE id = ?
`

// SQL Query to retrieve a movie by it's path name
const movieFindPathStmt = `
SELECT id, title, backdrop, poster, release_date, adult, genres, 
	homepage, imdb_id, overview, runtime, tagline, rating, filename
FROM movies WHERE filename = ?
`

// SQL Query to retrieve a all movies
const movieStmt = `
SELECT id, title, backdrop, poster, release_date, adult, genres, 
	homepage, imdb_id, overview, runtime, tagline, rating, filename
FROM movies
ORDER BY title ASC
`

// Returns the Movie with the given Id.
func GetMovie(id int64) (*Movie, error) {
	movie := Movie{}
	err := meddler.QueryRow(db, &movie, movieFindIdStmt, id)
	return &movie, err
}

func GetMovieByFilename(path string) (*Movie, error) {
	movie := Movie{}
	err := meddler.QueryRow(db, &movie, movieFindPathStmt, path)
	return &movie, err
}

// Saves a Movie.
func SaveMovie(movie *Movie) error {
	if movie.Id == 0 {
		movie.Created = time.Now().UTC()
	}
	movie.Updated = time.Now().UTC()
	return meddler.Save(db, movieTable, movie)
}

// Deletes an existing Movie.
func DeleteMovie(id int64) error {
	db.Exec("DELETE FROM movies WHERE id = ?", id)
	return nil
}

// Returns a list of all Movies
func ListMovies() ([]*Movie, error) {
	var movies []*Movie
	err := meddler.QueryAll(db, &movies, movieStmt)
	return movies, err
}

// Returns a list of Movies within the specified
// range (for pagination purposes).
func ListMoviesRange(limit, offset int) ([]*Movie, error) {
	var movies []*Movie
	err := meddler.QueryAll(db, &movies, movieStmt)
	return movies, err
}
