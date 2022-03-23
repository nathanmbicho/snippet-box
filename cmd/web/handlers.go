package main

import (
	"fmt"
	"github.com/nathanmbicho/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

//home
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//get latest snippets from SnippetModel.Latest
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.gohtml", &templateData{
		Snippets: s,
	})
}

/**
show snippet
define showSnippet as method against *application
*/

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	//check if id passed is valid
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// fetch data using SnippetModel.get and if error
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.gohtml", &templateData{
		Snippet: s,
	})

}

//createSnippetForm
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet...."))
}

/**
create new snippet
define createSnippet as method against *application
*/
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	//insert dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	//call SnippetModel.Insert method and pass data to execute
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//redirect to created snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
