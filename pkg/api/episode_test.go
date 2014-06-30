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

func createTestEpisode() *Episode {
	return &Episode{
		Title:          "Jack Falls Down the Well",
		Overview:       "Bad things are always happening, but Lassie will save you",
		Director:       "Bill Clinton",
		Writer:         "Bob Dole",
		GuestStars:     "Wilma Flintstone, John Bon Jovi",
		SeasonId:       12,
		EpisodeNumber:  2,
		SeasonNumber:   12,
		AbsoluteNumber: "324",
		Language:       "English",
		Rating:         "5 Stars",
		ImdbId:         "abcd1234",
	}
}

func populateSampleEpisodes(t *testing.T, name string, count int) (*Series, *Season, []*Episode) {
	series := NewSeries(name)
	if err := database.SaveSeries(series); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	season, err := NewSeason(series.Id, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := database.SaveSeason(season); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	episodes := make([]*Episode, count)
	for i := 0; i < count; i++ {
		e := &Episode{
			Title:        fmt.Sprintf("Episode %d", i),
			SeasonId:     season.Id,
			SeasonNumber: uint64(season.SeasonNumber),
			SeriesId:     series.Id,
		}
		if err := database.SaveEpisode(e); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		episodes[i] = e
	}

	return series, season, episodes
}

func TestListEpisodes(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	var count int = 3
	series, season, _ := populateSampleEpisodes(t, "test 1", count)

	// Create another series with episodes to ensure we aren't returning all episodes
	populateSampleEpisodes(t, "test 2", 10)

	r, err := http.Get(fmt.Sprintf("%s/api/series/%d/seasons/%d/episodes", s.URL, series.Id, season.SeasonNumber))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response []Episode
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(response) != count {
		t.Errorf("expected %d results, got %d", count, len(response))
	}
}

func TestGetEpisode(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	var count int = 4
	populateSampleEpisodes(t, "test 1", count)

	s1, err := database.GetEpisode(3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r, err := http.Get(fmt.Sprintf("%s/api/episodes/%d", s.URL, s1.Id))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var response Episode
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

func TestCreateEpisode(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	episode := createTestEpisode()
	b, err := json.Marshal(episode)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	r, err := Put(s.URL+"/api/episodes", "application/json", strings.NewReader(fmt.Sprintf("%s", b)))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var response Episode
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if response.Id == 0 {
		t.Errorf("expected non-zero episode Id, got %v", response.Id)
	}

	if response.Title != episode.Title {
		t.Errorf("expected matching titles, got %v", response)
	}

	// Check that it is actually in the database
	_, err = database.GetEpisode(response.Id)
	if err != nil {
		t.Errorf("Failed to find created episode in database: %v", err)
	}
}

func TestDeleteEpisode(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	var count int = 4
	populateSampleEpisodes(t, "test 1", count)

	r, err := Delete(fmt.Sprintf("%s/api/episodes/%d", s.URL, 1))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("expected %d status code, got %d", http.StatusOK, r.StatusCode)
	}

	var rm Episode
	err = json.NewDecoder(r.Body).Decode(&rm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if 1 != rm.Id {
		t.Errorf("expected returned item id %d, got id %d", 1, rm.Id)
	}

	// Double check this doesn't exist in the database
	if episode, err := database.GetEpisode(1); err == nil {
		t.Errorf("expected error fetching deleted episode, got %v", episode)
	}
}
