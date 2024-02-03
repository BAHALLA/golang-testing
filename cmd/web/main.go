package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {

	// config
	app := application{}

	// get session
	app.Session = getSession()

	// routes
	mux := app.routes()

	// start the server
	log.Println("Starting server at 8888 ... ")

	err := http.ListenAndServe(":8888", mux)

	if err != nil {
		log.Fatal(err)
	}

}
