package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestAppRoutes(t *testing.T) {

	var registred = []struct {
		route  string
		method string
	}{
		{route: "/", method: "Get"},
		{route: "/login", method: "Post"},
		{route: "/static/*", method: "Get"},
	}

	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registred {

		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registred", route.route)
		}
	}
}

func routeExists(testRoute string, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(testRoute, route) && strings.EqualFold(testMethod, testMethod) {
			found = true
		}
		return nil
	})
	return found
}
