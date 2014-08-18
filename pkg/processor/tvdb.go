package processor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rharter/go-guessit"
	"github.com/rharter/go-tmdb"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

func init() {
	con = tmdb.Open(TMDB_API_KEY)
}

type FetchSeriesMetadataTask struct {
	Element *Element
}

// FetchMetadataTask interface
func (t *FetchSeriesMetadataTask) Fetch() error {

	// See if we already have a element for this file
	element, err := database.GetElementByFile(t.Element.File)
	if err != nil {
		element = t.Element
	}
	element.ParentId = t.Element.ParentId

	meta, err := guessit.Guess(element.File)
	if err != nil {
		log.Printf("Failed to guess element info: %+v", err)
		return err
	}

	log.Printf("Searching for existing series %s", meta.Series)
	series, err := t.getOrCreateElementByTitle(meta.Series, "series", nil)
	if err != nil {
		log.Printf("Failed to find or create series[%s]: %v", meta.Series, err)
		return err
	}

	// Get the Series info
	log.Printf("Searching TMDB for series %s", series.Title)
	results, err := con.SearchTvShows(series.Title)
	if err != nil {
		log.Printf("Failed to get extra metadata for series %s: %v", results, err)
	} else {
		if len(results) > 0 {
			r := results[0]
			series.Title = r.Name
			series.Description = r.Overview
			series.Poster = r.PosterPath
			series.Background = r.BackdropPath
			series.RemoteId = string(r.Id)
			err = database.SaveElement(series)
			if err != nil {
				log.Printf("Failed to save series metadata %s: %v", series.Title, err)
			}
		}
	}

	log.Printf("Searching for existing season %d", meta.Season)
	season, err := t.getOrCreateElementByTitle(fmt.Sprintf("Season %d", meta.Season), "season", series)
	if err != nil {
		log.Printf("Failed to find or create season %d: %v", meta.Season, err)
		return err
	}

	// Get the season info
	log.Printf("Searching TMDB for season %d", meta.Season)
	var s *tmdb.Season
	id, err := strconv.ParseInt(series.RemoteId, 10, 64)
	if err != nil {
		log.Print("Failed to find series on tmdb.")
		season.Title = fmt.Sprintf("Season %d", meta.Season)
	} else {
		s, err = con.GetSeasonByNumber(id, meta.Season)
		if err != nil {
			log.Printf("Failed to get extra metadata for season %d: %v", meta.Season, err)
		} else {
			season.RemoteId = string(s.Id)
			season.Title = s.Name
			season.Description = s.Overview
			season.Poster = s.PosterPath
		}
	}
	err = database.SaveElement(season)
	if err != nil {
		log.Printf("Failed to save season metadata: %v", err)
	}

	log.Printf("Searcing for existing episode %s", meta.Title)
	episode, err := t.getOrCreateElementByTitle(meta.Title, "episode", season)
	if err != nil {
		log.Printf("Failed to find or create episode %s: %v", meta.Title, err)
		return err
	}

	var e *tmdb.Episode
	if s != nil {
		for _, val := range s.Episodes {
			if val.EpisodeNumber == meta.EpisodeNumber {
				e = val
				break
			}
		}
	} else {
		// TODO Search tmdb Directly by series id, season num and ep num
	}
	if e == nil {
		log.Printf("Failed to find episode number %d", meta.EpisodeNumber)
		return nil
	}

	episode.RemoteId = string(e.Id)
	episode.Title = e.Name
	episode.Description = e.Overview
	episode.Poster = e.StillPath
	err = database.SaveElement(episode)
	if err != nil {
		return err
	}

	return nil
}

func (task *FetchSeriesMetadataTask) getOrCreateElementByTitle(title string, t string, parent *Element) (*Element, error) {
	element, err := database.GetElementByTitle(title)
	if err != nil {
		if parent != nil {
			element = NewElement("", parent.Id, t)
		} else {
			element = NewElement("", -1, t)
		}
		element.Title = title
		err := database.SaveElement(element)
		if err != nil {
			return nil, err
		}
	}
	return element, nil
}
