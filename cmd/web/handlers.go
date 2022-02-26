package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//home
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//check if path and return not found
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	//check if files exists
	if err != nil {
		//access error to application logger
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	//check if route exists
	if err != nil {
		app.serverError(w, err)
		return
	}
}

/**
show snippet
define showSnippet as method against *application
*/

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	//check if id passed is valid
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display snippet with id :- %d..", id)
}

/**
create new snippet
define createSnippet as method against *application
*/
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	//check create method
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create new snippet..."))
}
