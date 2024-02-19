package main

import (
	"log"
	"net/http"
)

func main() {

	mux := setupHandlers()

	server := &http.Server{Addr: ":8080", Handler: mux}

	log.Printf("Listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Panicf("Server is not started %v", err)
	}

}
