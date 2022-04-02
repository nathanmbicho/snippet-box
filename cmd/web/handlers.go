package main

import (
	"fmt"
	"github.com/nathanmbicho/snippetbox/pkg/forms"
	"github.com/nathanmbicho/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

//home
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//get latest snippets from SnippetModel.Latest
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.gohtml", &templateData{
		Snippets: s,
	})
}

/**
show snippet
define showSnippet as method against *application
*/

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	//check if id passed is valid
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// fetch data using SnippetModel.get and if error
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.gohtml", &templateData{
		Snippet: s,
	})

}

//createSnippetForm
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.gohtml", &templateData{
		Form: forms.New(nil),
	})
}

/**
create new snippet
define createSnippet as method against *application
*/
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	//r.ParseForm to add POST form data request body to r.PostForm map
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//forms.Form struct to get relevant posted form data and use the validation methods to check the content
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	//check form.Valid if it returns len of errors or isn't valid, then redisplay the template passing in form.Form data
	if !form.Valid() {
		app.render(w, r, "create.page.gohtml", &templateData{
			Form: form,
		})
		return
	}

	//call SnippetModel.Insert method and pass data to execute
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	//use Put() method pass session success flash message
	app.session.Put(r, "flash", "Snippet successfully created!")

	//redirect to created snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

//-------AUTH FUNCTIONS HANDLERS-------//

//signupUserForm -Display the user register form...
func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.gohtml", &templateData{
		Form: forms.New(nil),
	})
}

//signupUser -create a new user
func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	//parse form data
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//validate form data
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchPattern("email", forms.EmailRx)
	form.MinLength("password", 8)

	//if errors found send back and redisplay
	if !form.Valid() {
		app.render(w, r, "signup.page.gohtml", &templateData{Form: form})
		return
	}

	//try to create new user record
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))

	//if email already exists and or if an error message
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Email address is already in use.")
		app.render(w, r, "signup.page.gohtml", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	//confirmation flash message if signup was successful
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")

	//redirect to login page
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

//loginUserForm - display the user login form
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

//loginUser - authenticate and login the user
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

//logoutUser - logout the user
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
