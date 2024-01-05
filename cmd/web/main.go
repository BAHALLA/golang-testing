package main

import (
	"log"
	"net/http"
)

type app struct{}

func main() {

	// config
	app := app{}

	// routes
	mux := app.routes()

	// start the server
	log.Println("Starting server at 8888 ... ")

	err := http.ListenAndServe(":8888", mux)

	if err != nil {
		log.Fatal(err)
	}

}
