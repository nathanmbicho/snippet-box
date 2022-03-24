package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler { //return http.Handler instead of a *http.ServeMux

	//create a middleware chain containing standard middleware which will be used for every request our application receives
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	//new middleware chain containing dynamic application routes containing session middleware
	dynamicMiddleware := alice.New(app.session.Enable)

	//serverMux
	mux := pat.New()

	//register routes
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	//serving static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Wrap the existing chain with the recoverPanic middleware.
	return standardMiddleware.Then(mux)
}
