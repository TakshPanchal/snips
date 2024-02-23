package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// TODO: Create a New Method for Application
type Application struct {
	infoLogger, errLogger *log.Logger
	db                    *sql.DB
}

func (a *Application) home(w http.ResponseWriter, r *http.Request) {
	a.infoLogger.Println(r.URL)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	homeFiles := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	tmpl, err := template.ParseFiles(homeFiles...)

	if err != nil {
		a.errLogger.Println("Error: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)

	if err != nil {
		a.errLogger.Println("Error: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (a *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryId)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "/snippet/view endpoint for id %d", id)
}

// snippetCreate serve /snippet/create endpoint to create a single snippet
// only POST request is allowed
func (a *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check for POST Method
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
		return
	}

	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, now(), now() + INTERVAL '1 days') RETURNING id;`

	_, err := a.db.Exec(query, "taksh", "taksh content")

	if err != nil {
		a.errLogger.Println("Error: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "/snippet/create endpoint")
}
