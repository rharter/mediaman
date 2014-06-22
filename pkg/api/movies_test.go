package api

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/pat"
	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

const TEST_DATABASE string = "/tmp/mediaman_test.sqlite"

func testMovie() *Movie {
	return &Movie{
		Title:        "Title",
		BackdropPath: "BackdropPath",
		PosterPath:   "PosterPath",
		ReleaseDate:  time.Now(),
		Adult:        true,
		Genres:       "Genres",
		Homepage:     "Homepage",
		ImdbID:       "ImdbID",
		Overview:     "Overview",
		Runtime:      140,
		Tagline:      "Tagline",
		UserRating:   1.5,
		Filename:     "Filename",
	}
}

func createTestServer() *httptest.Server {
	// Initialize the test database
	database.Init("sqlite3", TEST_DATABASE)

	// Create the handlers
	m := pat.New()
	AddHandlers(m, "/api")
	return httptest.NewServer(m)
}

func TestCreateMovie(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	movie := testMovie()
	b, err := json.Marshal(movie)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	r, err := Put(s.URL+"/api/movies", "application/json", strings.NewReader(fmt.Sprintf("%s", b)))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var returnedMovie Movie
	err = json.NewDecoder(r.Body).Decode(&returnedMovie)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if returnedMovie.ID == 0 {
		t.Errorf("expected non-zero movie ID, got %v", returnedMovie.ID)
	}

	if returnedMovie.Title != movie.Title {
		t.Errorf("expected matching titles, got %v", returnedMovie)
	}

	// Check that it is actually in the database
	_, err = database.GetMovie(returnedMovie.ID)
	if err != nil {
		t.Errorf("Failed to find created movie in database: %v", err)
	}
}
