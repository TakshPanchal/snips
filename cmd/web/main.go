package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	infoLogger := log.New(os.Stdin, "DEBUG:\t", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Llongfile)

	connStr := "user=server password='1234' dbname=snippets sslmode=disable"
	db, err := setupDB(connStr)
	if err != nil {
		errLogger.Panicf("Database connection failed: %v", err)
	}

	app := &Application{infoLogger: infoLogger, errLogger: errLogger, db: db}

	mux := setupHandlers(app)

	server := &http.Server{Addr: ":8080", Handler: mux, ErrorLog: errLogger}

	infoLogger.Printf("Listening on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		errLogger.Panicf("Server is not started %v", err)
	}

}

// setupDB creates a connection to the database for given dsn
func setupDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
