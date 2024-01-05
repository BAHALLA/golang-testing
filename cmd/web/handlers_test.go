package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppHandlers(t *testing.T) {

	var tests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{name: "home", url: "/", expectedStatusCode: http.StatusOK},
		{name: "404", url: "/not-exist", expectedStatusCode: http.StatusNotFound},
	}

	var app app

	routes := app.routes()

	ts := httptest.NewTLSServer(routes)

	for _, e := range tests {

		res, err := ts.Client().Get(ts.URL + e.url)

		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if res.StatusCode != e.expectedStatusCode {
			t.Errorf("%s: expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
		}
	}

}
