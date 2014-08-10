package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
	"github.com/rharter/mediaman/pkg/processor"
)

// Display a library list
func LibraryList(w http.ResponseWriter, r *http.Request) error {
	libraries, err := database.ListLibraries()
	if err != nil {
		return err
	}

	data := struct {
		Libraries []*Library `json:"libraries"`
	}{libraries}

	return RenderTemplate(w, r, "list_libraries.html", &data)
}

// Display a video list for a library
func LibraryShow(w http.ResponseWriter, r *http.Request) error {
	idstr := r.URL.Query().Get(":id")

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return err
	}

	lib, err := database.GetLibrary(id)
	if err != nil {
		return err
	}

	switch lib.Type {
	case "movies":
		return movieList(lib, w, r)
	default:
		return errors.New(fmt.Sprintf("No handler to process library of type %s", lib.Type))
	}
}

// For displaying movies, the displayed movies will always be the direct
// decendant of the library's root directory.
func movieList(l *Library, w http.ResponseWriter, r *http.Request) error {
	vids, err := database.GetElementsForParent(l.RootId)
	if err != nil {
		return err
	}

	data := struct {
		Elements []*Element
	}{vids}

	return RenderTemplate(w, r, "movies_list.html", data)
}

func LibraryNew(w http.ResponseWriter, r *http.Request) error {
	return RenderTemplate(w, r, "libraries_create.html", nil)
}

func LibraryCreate(w http.ResponseWriter, r *http.Request) error {
	library := &Library{
		Name: r.FormValue("name"),
		Type: r.FormValue("type"),
		Root: &Element{
			File: r.FormValue("path"),
		},
	}

	err := database.SaveLibrary(library)
	if err != nil {
		return err
	}

	http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)

	return nil
}

func LibraryProcess(w http.ResponseWriter, r *http.Request) error {
	idstr := r.URL.Query().Get(":id")

	log.Printf("Handling request to precess library: %s", idstr)

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return err
	}

	library, err := database.GetLibrary(id)
	if err != nil {
		return err
	}

	processor.ProcessLibrary(library)

	return RenderText(w, "Processing", 200)
}
