package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/takshpanchal/snips/internal/models"
)

// TODO: Create a New Method for Application
type Application struct {
	infoLogger, errLogger *log.Logger
	snippetModel          *models.SnippetModel
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

// snippetView is http handler for "/snippet/view" endpoint
func (a *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryId)

	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	a.infoLogger.Printf("id: %d", id)

	snip, err := a.snippetModel.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)
			return
		} else {
			a.serverError(w, err)
			return
		}
	}

	fmt.Fprintf(w, "Found Snippet: %+v", snip)
}

// snippetCreate serve /snippet/create endpoint to create a single snippet
// only POST request is allowed
func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check for POST Method
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	snip := models.Snippet{
		Title:   "Simple Snippet Title",
		Content: "Lorem ipsum content",
	}
	id, err := app.snippetModel.Insert(snip)

	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "Snippet is created with an id of %d", id)
}
