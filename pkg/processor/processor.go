package processor

import (
	"database/sql"
	"log"
	"mime"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
)

func ProcessLibrary(library *Library) (err error) {
	log.Printf("Processing library for path: %s", library.Root.File)
	queue := Start(runtime.NumCPU()*2 + 1)

	// Cleanup the library first.
	cleanup(library.Root, library)

	var chann chan FetchMetadataTask
	switch library.Type {
	case "movies":
		chann = processMovieDir(library)
	case "series":
		chann = processTvShowDir(library)
	}

	for task := range chann {
		queue.Add(task)
	}
	return nil
}

// Recursively steps through all children to remove
// elements whose file no longer exists.
func cleanup(e *Element, l *Library) error {
	if l == nil || (l.Root != e && e.File != "") {
		if _, err := os.Stat(e.File); os.IsNotExist(err) {
			log.Printf("Removing dangling reference to file: %s", e.File)
			// The delete operation will remove all children, so we
			// don't have to worry about continuing
			return database.DeleteElementCascade(e.Id)
		}
	}

	// Get the children
	es, err := database.GetElementsForParent(e.Id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, el := range es {
		err = cleanup(el, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func processMovieDir(l *Library) chan FetchMetadataTask {
	chann := make(chan FetchMetadataTask)
	go func() {
		filepath.Walk(l.Root.File, func(path string, info os.FileInfo, _ error) (err error) {
			if info != nil && !info.IsDir() {
				ext := filepath.Ext(path)
				mimetype := mime.TypeByExtension(ext)
				if strings.HasPrefix(mimetype, "video") {
					task := FetchMovieMetadataTask{
						Element: NewElement(path, l.Root.Id, "movie"),
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

func processTvShowDir(l *Library) chan FetchMetadataTask {
	chann := make(chan FetchMetadataTask)
	go func() {
		filepath.Walk(l.Root.File, func(path string, info os.FileInfo, _ error) (err error) {
			if info != nil && !info.IsDir() {
				ext := filepath.Ext(path)
				mimetype := mime.TypeByExtension(ext)
				if strings.HasPrefix(mimetype, "video") {
					task := FetchSeriesMetadataTask{
						Element: NewElement(path, l.Root.Id, "series"),
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
