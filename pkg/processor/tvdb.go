package processor

import (
	"fmt"
	"log"

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

	var query string
	var meta *guessit.GuessResult
	if element.Title != "" {
		query = element.Title
	} else {
		// Try to get the file name info from GuessIt
		meta, err = guessit.Guess(element.File)
		if err != nil {
			log.Printf("Failed to guess element info: %+v", err)
			return err
		}
	}
	log.Printf("Fetching metadata for %s", query)

	series, err := t.getOrCreateElementByTitle(meta.Series, "series", nil)
	if err != nil {
		log.Printf("Failed to find or create series[%s]: %v", meta.Series, err)
		return err
	}

	// Get the Series info
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

	season, err := t.getOrCreateElementByTitle(fmt.Sprintf("Season %d", meta.Season), "season", series)
	if err != nil {
		log.Printf("Failed to find or create season %d: %v", meta.Season, err)
		return err
	}

	// Get the season info
	s, err := con.GetSeasonByNumber(results[0].Id, meta.Season)
	if err != nil {
		log.Printf("Failed to get extra metadata for season %d: %v", meta.Season, err)
	} else {
		season.RemoteId = string(s.Id)
		season.Title = s.Name
		season.Description = s.Overview
		season.Poster = s.PosterPath
		err = database.SaveElement(season)
		if err != nil {
			log.Printf("Failed to save season metadata: %v", err)
		}
	}

	episode, err := t.getOrCreateElementByTitle(meta.Title, "episode", season)
	if err != nil {
		log.Printf("Failed to find or create episode %s: %v", meta.Title, err)
		return err
	}

	var e *tmdb.Episode
	for _, val := range s.Episodes {
		if val.EpisodeNumber == meta.EpisodeNumber {
			e = val
			break
		}
	}
	if e == nil {
		log.Printf("Failed to find episode number %d", meta.EpisodeNumber)
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
		element.Title = t
		err := database.SaveElement(element)
		if err != nil {
			return nil, err
		}
	}
	return element, err
}
