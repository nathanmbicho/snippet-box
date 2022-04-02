package main

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/nathanmbicho/snippetbox/pkg/models"
	"net/http"
	"runtime/debug"
	"time"
)

/**
serverError to handle error 500
*/
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	err = app.errorLog.Output(2, trace)
	if err != nil {
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/**
clientError to send specific status code error description
e.g -error 400 - bad request
*/
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

/**
notFound to handle error 404
*/
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

//addDefaultData helper
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.CSRFToken = nosurf.Token(r) //add CSRF token to template data and make it available each time when rendering the page form
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash") // automate display of any flash message on any page
	td.AuthenticatedUser = app.authenticatedUser(r)
	return td

}

//render template cache error
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	//check if template file exists, else throw serverError
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	// Initialize a new buffer
	buf := new(bytes.Buffer)

	//execute the template set passing in any dynamic data
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return
	}
}

//authenticatedUser -return ID of the current user from the session - 0 if not authenticated
func (app *application) authenticatedUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}
