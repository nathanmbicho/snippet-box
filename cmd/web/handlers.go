package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//home
func home(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
}

//show snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	//check if id passed is valid
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display snippet with id :- %d..", id)
}

//create new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	//check create method
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Write([]byte("Create new snippet..."))
}
