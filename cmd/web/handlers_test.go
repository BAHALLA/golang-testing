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

	req, _ := http.NewRequest("GET", "/", nil)

	req = addContextAndSessionToRequest(req, app)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Test Home return wrong http status, Got %d instead od 200", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)

	if !strings.Contains(string(body), ` <small> From session:`) {
		t.Errorf("did not found correct text in the body")
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
