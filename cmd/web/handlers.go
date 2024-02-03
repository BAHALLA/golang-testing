package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var td = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		td["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "this page was hit on "+time.Now().UTC().String())
	}

	_ = app.render(w, r, "home.page.gohtml", &TemplateData{Data: td})
}

// TemplateData define template data as a map
type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, templ string, data *TemplateData) error {

	parsedTemplate, err := template.ParseFiles(path.Join("./templates/"+templ), path.Join("./templates/", "base.layout.gohtml"))

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

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	form := NewForm(r.PostForm)

	form.Required("email", "password")

	if !form.Valid() {
		http.Error(w, "failed validation", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)
	fmt.Fprint(w, email)

}
