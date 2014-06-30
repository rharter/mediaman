package processor

import (
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/rharter/go-guessit"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

func ProcessLibrary(library *Library) (err error) {
	log.Printf("Processing library for path: %s", library.Path)
	queue := Start(20)
	chann := processDir(library.Path)
	for msg := range chann {
		filename := filepath.Base(msg)
		meta, err := guessit.Guess(filename)
		if err != nil {
			return err
		}

		switch meta.Type {
		case "movie":
			//processMovie(msg, queue)
		case "episode":
			processEpisode(msg, meta, queue)
		}
	}
	return nil
}

func processMovie(path string, queue *Queue) error {
	movie, _ := database.GetMovieByFilename(path)
	if movie == nil || movie.Filename == "" {
		movie, _ = NewMovie(path)
	}

	err := database.SaveMovie(movie)
	if err != nil {
		log.Printf("Error saving movie: %+v", err)
		return err
	}

	queue.Add(&FetchMetadataTask{Movie: movie})
	return nil
}

func processEpisode(path string, meta *guessit.GuessResult, queue *Queue) error {
	episode, _ := database.GetEpisodeByFilename(path)
	var fetchMetadataTask *FetchMetadataTask
	if episode == nil || episode.Filename == "" {
		s, err := database.GetSeriesByTitle(meta.Series)
		if err != nil {
			// Create and save a placeholder series
			s = &Series{
				Title: meta.Series,
			}
			err = database.SaveSeries(s)
			if err != nil {
				return err
			}

			// Queue a metadata refresh of the series data
			fetchMetadataTask = &FetchMetadataTask{Series: s}
		}

		episode = &Episode{
			Filename:      path,
			Title:         meta.Title,
			EpisodeNumber: uint64(meta.EpisodeNumber),
			SeasonNumber:  uint64(meta.Season),
			SeriesId:      s.Id,
			TvdbSeriesId:  s.SeriesId,
		}
		err = database.SaveEpisode(episode)
		if err != nil {
			log.Println("Error saving movie")
		}
	}

	if fetchMetadataTask != nil {
		queue.Add(fetchMetadataTask)
	} else {
		queue.Add(&FetchMetadataTask{Episode: episode})
	}

	return nil
}

func processDir(path string) chan string {
	chann := make(chan string)
	go func() {
		filepath.Walk(path, func(path string, info os.FileInfo, _ error) (err error) {
			if !info.IsDir() {
				ext := filepath.Ext(path)
				mimetype := mime.TypeByExtension(ext)
				if strings.HasPrefix(mimetype, "video") {
					chann <- path
				}
			}
			return
		})
		defer close(chann)
	}()
	return chann
}
