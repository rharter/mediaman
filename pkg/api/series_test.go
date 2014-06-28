package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

func createTestSeries() *Series {
	return &Series{
		Language: "english",
		Title:    "Game of Thrones",
		Overview: "Seven noble families fight for control of the mythical land of Westeros. Friction between the houses leads to full-scale war. All while a very ancient evil awakens in the farthest north. Amidst the war, a neglected military order of misfits, the Night's Watch, is all that stands between the realms of men and icy horrors beyond.",
		Banner:   "http://thetvdb.com/banners/graphical/121361-g22.jpg",
		ImdbId:   "1234abcd",
		SeriesId: 121361,
	}
}

func populateSampleSeries(t *testing.T, count int) {
	for pos := 0; pos < count; pos++ {
		s := &Series{Title: fmt.Sprintf("Series %d", pos)}
		if err := database.SaveSeries(s); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestListSeries_ListsAllSeries(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	var count int = 3
	populateSampleSeries(t, count)

	r, err := http.Get(s.URL + "/api/series")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response []Series
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(response) != count {
		t.Errorf("expected %d results, got %d", count, len(response))
	}
}

func TestGetSeries_ReturnsCorrectSeries(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	var count int = 4
	populateSampleSeries(t, count)

	s1, err := database.GetSeries(3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r, err := http.Get(fmt.Sprintf("%s/api/series/%d", s.URL, s1.Id))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response Series
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if s1.Id != response.Id {
		t.Errorf("expected id %d, got %d", s1.Id, response.Id)
	}

	if s1.Title != response.Title {
		t.Errorf("expected title %q, got %q", s1.Title, response.Title)
	}
}

func TestCreateSeries(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	series := createTestSeries()
	b, err := json.Marshal(series)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	r, err := Put(s.URL+"/api/series", "application/json", strings.NewReader(fmt.Sprintf("%s", b)))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var response Series
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if response.Id == 0 {
		t.Errorf("expected non-zero series Id, got %v", response.Id)
	}

	if response.Title != series.Title {
		t.Errorf("expected matching titles, got %v", response)
	}

	// Check that it is actually in the database
	_, err = database.GetSeries(response.Id)
	if err != nil {
		t.Errorf("Failed to find created series in database: %v", err)
	}
}
