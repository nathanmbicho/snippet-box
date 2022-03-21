package main

import "net/http"

func (app *application) routes() http.Handler { //return http.Handler instead of a *http.ServeMux

	//serverMux
	mux := http.NewServeMux()

	//register routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	//serving static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))
}
