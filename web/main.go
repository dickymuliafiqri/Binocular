package web

import (
	"net/http"
	"os"

	"github.com/dickymuliafiqri/Binocular/web/subfinder"
	"github.com/gorilla/handlers"
	gorillaMux "github.com/gorilla/mux"
)

func StartWebServer() error {
	mux := gorillaMux.NewRouter()

	mux.HandleFunc("/", subfinder.SubfinderHandler).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)
	return http.ListenAndServe(":8080", loggedRouter)
}
