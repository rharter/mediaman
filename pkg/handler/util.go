package handler

import (
	"net/http"
	"encoding/json"
	"strings"

	"github.com/rharter/mediaman/pkg/template"
)

// -------------------------------------------------------
// Rendering Functions

// Renders the named template for the specified data type 
// and write the output to the http.ResponseWriter
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) error {
	if strings.HasPrefix(r.URL.Path, "/api") {
		w.Header().Add("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(data)
	} else {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		return template.ExecuteTemplate(w, name, data)
	}
}

// RenderText write the plain text string to the http.ResponseWriter
func RenderText(w http.ResponseWriter, text string, code int) error {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(text))
	return nil
}