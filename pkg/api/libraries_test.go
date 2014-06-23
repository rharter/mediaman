package api

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

func testLibrary(name, path string) *Library {
	return &Library{
		Name: name,
		Path: path,
	}
}

func TestListLibraries(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	l1 := testLibrary("l1", "/tmp/l1")
	err := database.SaveLibrary(l1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	l2 := testLibrary("l2", "/tmp/l2")
	err = database.SaveLibrary(l2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r, err := http.Get(s.URL + "/api/libraries")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var libs []Library
	err = json.NewDecoder(r.Body).Decode(&libs)
	if err != nil {
		t.Errorf("unparsable response: %v", err)
	}

	if len(libs) != 2 {
		t.Errorf("expected 2 libraries, got %d", len(libs))
	}
}
