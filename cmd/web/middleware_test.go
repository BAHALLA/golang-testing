package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppAddIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emtpryAddr  bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.25.26.1", "", false},
		{"", "", "hello:world", false},
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		val := r.Context().Value(contextUserKey)

		if val == nil {
			t.Errorf("%s not present", contextUserKey)
		}

		ip, ok := val.(string)

		if !ok {
			t.Error("not a string")
		}
		t.Log(ip)
	})

	for _, e := range tests {

		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://localhost", nil)

		if e.emtpryAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func TestAppIpFromContext(t *testing.T) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextUserKey, "somevalue")

	ip := app.ipFromContext(ctx)

	if !strings.EqualFold(ip, "somevalue") {
		t.Errorf("Expected somevalue but got %s", ip)
	}
}
