package handler

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/rharter/mediaman/pkg/database"
	. "github.com/rharter/mediaman/pkg/model"
	"github.com/rharter/mediaman/pkg/transcoder"
)

// Display a element list for a library
func ElementShow(w http.ResponseWriter, r *http.Request) error {
	idstr := r.URL.Query().Get(":id")

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return err
	}

	// For now, assume all elements are movies
	el, err := database.GetElement(id)
	if err != nil {
		return err
	}

	data := struct {
		Element  *Element
		MimeType string
	}{
		Element:  el,
		MimeType: mime.TypeByExtension(filepath.Ext(el.File)),
	}

	return RenderTemplate(w, r, "movies_show.html", data)
}

// GET /elements/:id/element
func ElementPlay(w http.ResponseWriter, r *http.Request) {
	id, err := elementIdFromRequest(r)
	if err != nil {
		return
	}

	element, err := database.GetElement(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, element.File)
}

func elementIdFromRequest(r *http.Request) (int64, error) {
	idstr := r.URL.Query().Get(":id")
	return strconv.ParseInt(idstr, 10, 64)
}

// GET /elements/:id/transcode
func ElementTranscode(w http.ResponseWriter, r *http.Request) {
	id, err := elementIdFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	element, err := database.GetElement(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	log.Printf("Starting transcodeing session for file: %s", element.File)
	trans := transcoder.NewTranscoderSession(element.Id, "/tmp/element_parts", element.File)

	if err = trans.Open(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rel, err := filepath.Rel("/tmp/element_parts", trans.OutputFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := filepath.Join("/elements/parts/", rel)

	log.Printf("Serving output file: %s", p)
	http.Redirect(w, r, p, http.StatusMovedPermanently)
}
