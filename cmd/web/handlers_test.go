package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestAppHome(t *testing.T) {

	var tests = []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{"First visit", "", "<small> From session:"},
		{"second visit", "hello, world!", "<small> From session: hello, world!"},
	}

	for _, e := range tests {

		req, _ := http.NewRequest("GET", "/", nil)

		req = addContextAndSessionToRequest(req, app)

		_ = app.Session.Destroy(req.Context())

		if e.putInSession != "" {
			app.Session.Put(req.Context(), "test", e.putInSession)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.Home)

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Test Home return wrong http status, Got %d instead od 200", rr.Code)
		}

		body, _ := io.ReadAll(rr.Body)

		if !strings.Contains(string(body), e.expectedHTML) {
			t.Errorf("%s: did not find %s in the response body", e.name, e.expectedHTML)
		}
	}

}

func getContext(req *http.Request) context.Context {
	return context.WithValue(req.Context(), contextUserKey, "unknown")
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {

	req = req.WithContext(getContext(req))
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
