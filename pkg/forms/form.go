package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

//Form struct to hold form values and the errors
type Form struct {
	url.Values
	Errors errors
}

//New function to initialize a custom Form struct by taking the form data as a parameter
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required - check if input field is empty or null
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

//MaxLength - check maximum number of characters passed in the form field
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters).", d))
	}
}

//PermittedValues method to check if field value matches the specific permitted values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid.")
}

//Valid method to return true if no errors are found
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
