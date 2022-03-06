package main

import (
	"bytes"
	"fmt"
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
	td.CurrentYear = time.Now().Year()
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

	buf.WriteTo(w)
}
