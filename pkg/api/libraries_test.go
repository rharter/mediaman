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

func TestCreateLibrary(t *testing.T) {
	s := createTestServer()
	defer s.Close()
	defer os.Remove(TEST_DATABASE)

	l1 := testLibrary("l1", "/tmp/l1")

	b, err := json.Marshal(l1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	reader := strings.NewReader(fmt.Sprintf("%s", b))
	r, err := Put(s.URL+"/api/libraries", "application/json", reader)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var resp Library
	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.ID == 0 {
		t.Errorf("expected non-zero id, got %d", resp.ID)
	}

	l2, err := database.GetLibrary(resp.ID)
	if err != nil {
		t.Errorf("library not saved to db with id %d", resp.ID)
	}

	if l2.Name != l1.Name {
		t.Errorf("expected name %q, got %q", l1.Name, l2.Name)
	}

	if l2.Path != l1.Path {
		t.Errorf("expected path %q, got %q", l1.Path, l2.Path)
	}
}
