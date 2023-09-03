package api

import (
	"net/http"
	"os"

	"github.com/dickymuliafiqri/Binocular/api/subfinder"
	"github.com/gorilla/handlers"
	gorillaMux "github.com/gorilla/mux"
	"github.com/rs/cors"
)

var webRoot = os.Getenv("WEB_ROOT")
var c = cors.New(cors.Options{
	AllowedOrigins: []string{"*"},
})

func StartWebServer() error {
	// Local deployment
	if webRoot == "" {
		webRoot = "web/dist"
	}

	mux := gorillaMux.NewRouter()

	mux.HandleFunc("/subfinder", subfinder.SubfinderHandler).Methods("GET")
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot)))

	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)
	return http.ListenAndServe(":8080", c.Handler(loggedRouter))
}
