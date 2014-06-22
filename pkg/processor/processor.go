package processor

import (
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

func ProcessLibrary(library *Library) (err error) {
	log.Printf("Processing library for path: %s", library.Path)
	queue := Start(5)
	chann := processDir(library.Path)
	for msg := range chann {
		movie, _ := database.GetMovieByFilename(msg)

		if movie.Filename == "" {
			movie, err = NewMovie(msg)
			if err != nil {
				log.Printf("Error creating new movie: %v", err)
				return err
			}
		}

		err = database.SaveMovie(movie)
		if err != nil {
			log.Printf("Error saving movie: %+v", err)
			return err
		}

		queue.Add(&FetchMetadataTask{Movie: movie})
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
