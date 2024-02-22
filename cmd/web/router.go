package main

import (
	"log"
	"net/http"
)

func setupHandlers(app *Application) *http.ServeMux {
	mux := http.NewServeMux()

	const STATIC_DIR = "./ui/static/"

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	log.Printf("Starting static file-server from %s", STATIC_DIR)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR))))

	return mux
}
