package handler

import (
	"log"
	"net/http"
)

// ErrorHandler wraps the default http.HandleFunc to handl an
// error as the return value
type ErrorHandler func(w http.ResponseWriter, r *http.Request) error

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		log.Print(err)
	}
}
