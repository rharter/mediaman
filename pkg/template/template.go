package template

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/GeertJohan/go.rice"
)

// ErrTemplateNotFound indicates the requested template
// does not exist in the TemplateStore
var ErrTemplateNotFound = errors.New("Template Not Found")

// registry stores a map of Templates where the key
// is the template name and the value is the *template.Template
var registry = map[string]*template.Template{}

// ExecuteTemplate applies the template to the specified data object and
// writes the output to wr.
func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	templ, ok := registry[name]
	if !ok {
		return ErrTemplateNotFound
	}
	return templ.ExecuteTemplate(wr, "_", data)
}

// all templates are loaded on initialization
func init() {
	// location of templates
	box := rice.MustFindBox("pages")

	var files = []string{
		"libraries_create.html",
		"list_libraries.html",

		"movies_list.html",
		"movies_show.html",
		"404.html",
	}

	// extract the base templte as a string
	base, err := box.String("base.html")
	if err != nil {
		panic(err)
	}

	// loop through files and create templates
	for _, file := range files {
		// extract the template from the box
		page, err := box.String(file)
		if err != nil {
			panic(err)
		}

		// parse the template and then add to the global map
		baseParsed, err := template.New("_").Parse(base)
		if err != nil {
			panic(fmt.Errorf("Error parsing base.html template: %s", err))
		}
		pageParsed, err := baseParsed.Parse(page)
		if err != nil {
			panic(fmt.Errorf("Error parsing page template for %s: %s", file, err))
		}

		registry[file] = pageParsed
	}
}
