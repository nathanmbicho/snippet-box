package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
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
