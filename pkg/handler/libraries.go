package handler

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"

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

func LibraryCreate(w http.ResponseWriter, r *http.Request) error {
	var library Library
	err := json.NewDecoder(r.Body).Decode(&library)
	if err != nil {
		return err
	}

	err = database.SaveLibrary(&library)
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