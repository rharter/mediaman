package processor

import (
	"os"
	"log"
	"mime"
	"path/filepath"
	"strings"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/rharter/mediaman/pkg/database"
)

func ProcessLibrary(library *Library) (err error) {
	log.Printf("Processing library for path: %s", library.Path)
	queue := Start(5)
	chann := processDir(library.Path)
	for msg := range chann {
		movie, _ := database.GetMovieByPath(msg)

		if movie.Filename == "" {
			movie = NewMovie(msg)
		}

		err = database.SaveMovie(movie)
		if err != nil {
			log.Printf("Error saving movie: %+v", err)
			return err
		}

		queue.Add(&FetchMetadataTask{Movie:movie})
	}
	return nil
}

func processDir(path string) (chan string) {
	chann := make(chan string)
	go func() {
		filepath.Walk(path, func(path string, info os.FileInfo, _ error)(err error) {
			if !info.IsDir() {
				ext := filepath.Ext(path)
				mimetype := mime.TypeByExtension(ext)
				if (strings.HasPrefix(mimetype, "video")) {
					chann <- path
				}
			}
			return
		})
		defer close(chann)
	}()
	return chann
}