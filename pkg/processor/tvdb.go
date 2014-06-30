package processor

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/StalkR/imdb"
	"github.com/rharter/tvdb"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

var tvdbconn *tvdb.TVDB

const TVDB_BASE_IMAGE_PATH = "http://thetvdb.com"

func init() {
	tvdbconn = tvdb.Open("9965FF073DB50ECD")
}

func FetchMetadataForSeries(queue *Queue, series *Series) error {
	log.Printf("Fetching metadata for series: %q", series.Title)

	ss, err := tvdbconn.GetSeries(series.Title, "")
	if err != nil {
		log.Printf("Error querying TVDB for %s: %v", series.Title, err)
		return err
	}

	if len(ss) < 1 {
		log.Printf("No Series found for name: %s", series.Title)
		return nil
	}

	// Now that we have basic info, get hte extended info.
	s, err := tvdbconn.GetSeriesById(ss[0].Id, "")
	if err != nil {
		log.Printf("Error querying TVDB for %s: %v", series.Title, err)
		return err
	}

	// Use the top matching series for now
	series.Title = s.SeriesName
	series.Overview = s.Overview
	series.SeriesId = int64(s.Id)
	series.Language = s.Language
	series.ImdbId = s.ImdbId
	series.Banner = tvdbconn.GetImageUrl(s.Banner)
	series.FanArt = tvdbconn.GetImageUrl(s.FanArt)
	series.Poster = tvdbconn.GetImageUrl(s.Poster)

	log.Printf("Writing series: %+v", series)

	// Persist the series
	err = database.SaveSeries(series)
	if err != nil {
		return err
	}

	eps, err := database.ListEpisodesForSeries(series.SeriesId)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Failed to fetch episodes for season: %v", err)
	}

	if eps != nil {
		for _, ep := range eps {
			ep.TvdbSeriesId = series.SeriesId
			err = database.SaveEpisode(ep)
			if err != nil {
				return err
			}

			queue.Add(&FetchMetadataTask{Episode: ep})
		}
	}

	return nil
}

func FetchMetadataForEpisode(e *Episode) error {
	log.Printf("Fetching metadata for episode: %q", e.Title)

	if e.TvdbSeriesId == 0 {
		s, err := database.GetSeries(e.SeriesId)
		if err != nil {
			return err
		}
		e.TvdbSeriesId = s.SeriesId
		log.Printf("Adding series id to episode: %d", e.TvdbSeriesId)
	}

	b, err := tvdbconn.GetEpisodeBySeasonEp(int(e.TvdbSeriesId), int(e.SeasonNumber), int(e.EpisodeNumber), "en")
	if err != nil {
		log.Printf("Error querying TheTVDB for %s: %v", e.Title, err)
		return err
	}

	singleEpisode, err := tvdb.ParseSingleEpisode(b)
	if err != nil {
		log.Printf("Error parsing single episode: %v", err)
		return err
	}

	episode := singleEpisode.Episode
	e.Title = episode.EpisodeName
	e.Overview = episode.Overview
	e.ImdbId = episode.ImdbId
	e.Director = episode.Director
	e.Writer = episode.Writer
	e.SeasonId = int64(episode.SeasonId)
	e.GuestStars = episode.GuestStars
	e.AbsoluteNumber = episode.AbsoluteNumber

	// IMDB info
	title, err := imdb.NewTitle(http.DefaultClient, e.ImdbId)
	if err != nil {
		log.Printf("Failed to get extended info from imdb: %v", err)
	} else {
		e.Poster = title.Poster.ContentURL
	}

	err = database.SaveEpisode(e)
	if err != nil {
		return err
	}

	return nil
}
