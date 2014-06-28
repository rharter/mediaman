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
	queue := Start(5)
	chann := processDir(library.Path)
	for msg := range chann {
		meta, err := guessit.Guess(msg)
		if err != nil {
			return err
		}

		switch meta.Type {
		case "movie":
			processMovie(msg, queue)
		case "episode":
			processEpisode(msg, queue)
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

func processEpisode(path string, queue *Queue) error {
	log.Printf("Not actually processing episode: %s", path)
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
