package main

import (
	"net/url"
	"strings"
)

type errors map[string][]string

func (e errors) Get(field string) string {

	errorSlice := e[field]

	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

func (e errors) Add(field, msg string) {

	e[field] = append(e[field], msg)

}

// Form type define forms data ans erros
type Form struct {
	Data   url.Values
	Errors errors
}

// NewForm create new form
func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

// Has check if form has a field
func (f *Form) Has(field string) bool {

	x := f.Data.Get(field)

	return x != ""
}

// Required define field that are required in the form
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Check checks if form is valide
func (f *Form) Check(ok bool, key string, msg string) {

	if !ok {
		f.Errors.Add(key, msg)
	}
}

// Valid check if there is any error in the form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
