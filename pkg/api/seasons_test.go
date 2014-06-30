package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bmizerany/pat"
	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type SeasonTestSuite struct {
	server *httptest.Server
}

var _ = Suite(&SeasonTestSuite{})

// Before
func (s *SeasonTestSuite) SetUpTest(c *C) {
	// Initialize the test database
	database.Init("sqlite3", "/tmp/mediaman_test.sqlite")

	// Create the handlers
	m := pat.New()
	AddHandlers(m, "/api")
	s.server = httptest.NewServer(m)
}

// After
func (s *SeasonTestSuite) TearDownTest(c *C) {
	s.server.Close()
	os.Remove("/tmp/mediaman_test.sqlite")
}

func (s *SeasonTestSuite) TestListSeasons(c *C) {
	series, seasons := s.createSeasons(c, 4)

	r, err := http.Get(fmt.Sprintf("%s/api/series/%d/seasons", s.server.URL, series.Id))
	c.Assert(err, IsNil)

	var resp []Season
	err = json.NewDecoder(r.Body).Decode(&resp)
	c.Assert(err, IsNil)

	c.Assert(len(resp), Equals, len(seasons))
}

// Creates test seasons all belonging to a single series
func (s *SeasonTestSuite) createSeasons(c *C, count int) (*Series, []*Season) {
	series := NewSeries("test series")
	if err := database.SaveSeries(series); err != nil {
		c.Errorf("unexpected error: %v", err)
	}

	seasons := make([]*Season, count)
	for i := 0; i < count; i++ {
		s, err := NewSeason(series.Id, uint64(i+1))
		if err != nil {
			c.Errorf("unexpected error: %v", err)
		}
		if err = database.SaveSeason(s); err != nil {
			c.Errorf("unexpected error: %v", err)
		}
		seasons[i] = s
	}

	return series, seasons
}
