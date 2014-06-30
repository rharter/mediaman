package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/rharter/mediaman/pkg/api"
	"github.com/rharter/mediaman/pkg/database"
	"github.com/rharter/mediaman/pkg/handler"
)

var (
	// port the server will run on
	port string

	// database driver used to connect to the database
	driver string

	// driver specific connection information. In this
	// case, it should be the location of the SQLite file
	datasource string

	// commit sha for the current build
	version string

	// optional flags for tls listener
	sslcert string
	sslkey  string
)

func main() {

	// parse command line flags
	flag.StringVar(&port, "port", ":8080", "")
	flag.StringVar(&driver, "driver", "sqlite3", "")
	flag.StringVar(&datasource, "datasource", "mediaman.sqlite", "")
	flag.StringVar(&sslcert, "sslcert", "", "")
	flag.StringVar(&sslkey, "sslkey", "", "")
	flag.Parse()

	// validate the TLS arguments
	checkTLSFlags()

	// setup database and handlers
	if err := database.Init(driver, datasource); err != nil {
		log.Fatal("Can't initialize database: ", err)
	}
	setupHandlers()

	// debug
	log.Printf("starting mediaman version %s on port %s\n", version, port)

	// start webserver using HTTPS or HTTP
	if sslcert != "" && sslkey != "" {
		panic(http.ListenAndServeTLS(port, sslcert, sslkey, nil))
	} else {
		panic(http.ListenAndServe(port, nil))
	}
}

// Checks if the TLS flags were supplied correctly
func checkTLSFlags() {
	if sslcert != "" && sslkey == "" {
		log.Fatal("invalid configuration: -sslkey unspecified, but -sslcert was specified.")
	} else if sslcert == "" && sslkey != "" {
		log.Fatal("invalid configuration: -sslcert unspecified, but -sslkey was specified.")
	}
}

// Setup routes for serving dynamic content
func setupHandlers() {

	m := pat.New()

	api.AddHandlers(m, "/api")

	m.Get("/movies", handler.ErrorHandler(handler.MovieList))
	m.Get("/movies/:id", handler.ErrorHandler(handler.MovieShow))

	m.Get("/movies/:id/video", http.HandlerFunc(handler.MoviePlay))
	m.Get("/movies/:id/transcode", http.HandlerFunc(handler.MovieTranscode))
	m.Get("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("/tmp/videos"))))

	m.Get("/libraries", handler.ErrorHandler(handler.LibraryList))
	m.Post("/libraries", handler.ErrorHandler(handler.LibraryCreate))

	m.Get("/libraries/:id/process", handler.ErrorHandler(handler.LibraryProcess))

	m.Get("/series", handler.ErrorHandler(handler.SeriesList))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// standard header variables that should be set, for good measure.
		w.Header().Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		w.Header().Add("X-Frame-Options", "DENY")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-XSS-Protection", "1; mode=block")

		m.ServeHTTP(w, r)
	})
}
