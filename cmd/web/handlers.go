package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Application struct {
	infoLogger, errLogger *log.Logger
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

	a.infoLogger.Println("Create API endpoint")
	// check for POST Method
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "/snippet/create endpoint")
}
