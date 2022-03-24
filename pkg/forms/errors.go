package forms

//define type errors to hold form validation error messages and form field name as the key in the map
type errors map[string][]string

//Add method - to add error messages for a given field in the map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//Get method to retrieve the first error message for a given field in the map
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
