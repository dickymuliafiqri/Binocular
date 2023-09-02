package api

import (
	"net/http"
	"os"

	"github.com/dickymuliafiqri/Binocular/api/subfinder"
	"github.com/gorilla/handlers"
	gorillaMux "github.com/gorilla/mux"
	"github.com/rs/cors"
)

var c = cors.New(cors.Options{
	AllowedOrigins: []string{"*"},
})

func StartWebServer() error {
	mux := gorillaMux.NewRouter()

	mux.HandleFunc("/subfinder", subfinder.SubfinderHandler).Methods("GET")
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("web/dist/")))

	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)
	return http.ListenAndServe(":8080", c.Handler(loggedRouter))
}
