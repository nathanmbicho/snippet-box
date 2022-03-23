package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler { //return http.Handler instead of a *http.ServeMux

	//create a middleware chain containing standard middleware which will be used for every request our application receives
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//serverMux
	mux := http.NewServeMux()

	//register routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	//serving static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Wrap the existing chain with the recoverPanic middleware.
	return standardMiddleware.Then(mux)
}
