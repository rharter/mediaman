package processor

import (
	"log"
	"mime"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/rharter/mediaman/pkg/model"
)

func ProcessLibrary(library *Library) (err error) {
	log.Printf("Processing library for path: %s", library.Path)
	queue := Start(runtime.NumCPU()*2 + 1)

	var chann chan FetchMetadataTask
	switch library.Type {
	case "movies":
		chann = processMovieDir(library.Path)
	case "series":
		log.Fatalf("Series processing hasn't been implemented yet.")
	}

	for task := range chann {
		queue.Add(task)
	}
	return nil
}

func processMovieDir(path string) chan FetchMetadataTask {
	chann := make(chan FetchMetadataTask)
	go func() {
		filepath.Walk(path, func(path string, info os.FileInfo, _ error) (err error) {
			if !info.IsDir() {
				ext := filepath.Ext(path)
				mimetype := mime.TypeByExtension(ext)
				if strings.HasPrefix(mimetype, "video") {
					task := FetchMovieMetadataTask{
						Video: NewVideo(path),
					}
					chann <- &task
				}
			}
			return
		})
		defer close(chann)
	}()
	return chann
}
