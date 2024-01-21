package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {

	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
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

	data.IP = app.ipFromContext(r.Context())

	err = parsedTemplate.Execute(w, data)

	if err != nil {
		return err
	}
	return nil
}

func (app *app) Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)
	fmt.Fprint(w, email)

}
