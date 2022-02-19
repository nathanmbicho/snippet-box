package main

import (
	"log"
	"net/http"
)

//main function
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//serving static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
