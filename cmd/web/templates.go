package main

import (
	"github.com/nathanmbicho/snippetbox/pkg/models"
	"html/template"
	"net/url"
	"path/filepath"
	"time"
)

//define holding structure for any dynamic data we want to pass to HTML templates
type templateData struct {
	CurrentYear int
	FormData    url.Values        //hold form data
	FormErrors  map[string]string //hold form errors
	Snippet     *models.Snippet   // return single data
	Snippets    []*models.Snippet // return slice of data
}

//humanDate function to return readable date
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

//initialize a template.FuncMap object and store it in a global variable
var functions = template.FuncMap{
	"humanDate": humanDate,
}

//add catching to the template data
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	//initialize new map struct to act as the cache
	cache := map[string]*template.Template{}

	//return all file paths  with extension '.page.gohtml' as a slice using filepath.Glob function
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	//loop between the file pages
	for _, page := range pages {
		//extract file name from the full file path and assign it to the name variable
		name := filepath.Base(page)

		//parse the page template file in to the template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//parse 'layout' templates in our template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		//parse 'partial' templates in our template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}

		//add template set to the cache using the name of the page as the key like 'home.page.gohtml'
		cache[name] = ts
	}

	return cache, nil //return the map
}
