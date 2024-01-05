package main

import (
	"html/template"
	"net/http"
)

func (app *app) home(w http.ResponseWriter, r *http.Request) {

	_ = app.render(w, r, "home.page.gohtml", nil)
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *app) render(w http.ResponseWriter, r *http.Request, templ string, data *TemplateData) error {

	parsedTemplate, err := template.ParseFiles("./templates/" + templ)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	err = parsedTemplate.Execute(w, data)

	if err != nil {
		return err
	}
	return nil
}
