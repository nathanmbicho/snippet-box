package main

import "github.com/nathanmbicho/snippetbox/pkg/models"

//define holding structure for any dynamic data we want to pass to HTML templates
type templateData struct {
	Snippet  *models.Snippet   // return single data
	Snippets []*models.Snippet // return slice of data
}
