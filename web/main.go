package web

import (
	"net/http"
	"os"

	"github.com/dickymuliafiqri/Binocular/web/subfinder"
	"github.com/gorilla/handlers"
	gorillaMux "github.com/gorilla/mux"
)

func StartWebServer() {
	mux := gorillaMux.NewRouter()

	mux.HandleFunc("/", subfinder.SubfinderHandler).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)
	http.ListenAndServe(":8080", loggedRouter)
}
