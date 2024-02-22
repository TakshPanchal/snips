package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	infoLogger := log.New(os.Stdin, "DEBUG:\t", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Llongfile)

	app := &Application{infoLogger: infoLogger, errLogger: errLogger}

	mux := setupHandlers(app)

	server := &http.Server{Addr: ":8080", Handler: mux, ErrorLog: errLogger}

	infoLogger.Printf("Listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		errLogger.Panicf("Server is not started %v", err)
	}

}
