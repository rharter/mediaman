package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
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
		ImdbId:       "ImdbId",
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

func TestMovieList(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	m1 := &Movie{Title: "Fight Club", Filename: "m1"}
	if err := database.SaveMovie(m1); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	m2 := &Movie{Title: "Rio 2", Filename: "m2"}
	if err := database.SaveMovie(m2); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r, err := http.Get(s.URL + "/api/movies")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response []Movie
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(response) != 2 {
		t.Errorf("expected 2 results, got %d", len(response))
	}
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

	if returnedMovie.Id == 0 {
		t.Errorf("expected non-zero movie Id, got %v", returnedMovie.Id)
	}

	if returnedMovie.Title != movie.Title {
		t.Errorf("expected matching titles, got %v", returnedMovie)
	}

	// Check that it is actually in the database
	_, err = database.GetMovie(returnedMovie.Id)
	if err != nil {
		t.Errorf("Failed to find created movie in database: %v", err)
	}
}

func TestDeleteMovie(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	m := testMovie()
	if err := database.SaveMovie(m); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r, err := Delete(s.URL + "/api/movies/" + strconv.FormatInt(m.Id, 10))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("expected %d status code, got %d", http.StatusOK, r.StatusCode)
	}

	var rm Movie
	err = json.NewDecoder(r.Body).Decode(&rm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if m.Id != rm.Id {
		t.Errorf("expected returned item id %d, got id %d", m.Id, rm.Id)
	}

	// Double check this doesn't exist in the database
	if movie, err := database.GetMovie(m.Id); err == nil {
		t.Errorf("expected error fetching deleted movie, got %v", movie)
	}
}
