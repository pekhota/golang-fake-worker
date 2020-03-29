package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router  {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Use(jsonResponseMiddleware)

	return router
}

func jsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

var routes = Routes{
	Route{
		Name:        "JobCreator",
		Method:      http.MethodPost,
		Pattern:     "/api/jobs",
		HandlerFunc: makeJob,
	},
	Route{
		Name:        "JobsList",
		Method:      http.MethodGet,
		Pattern:     "/api/jobs",
		HandlerFunc: listJobs,
	},
	Route{
		Name:        "MainRoute",
		Method:      http.MethodGet,
		Pattern:     "/",
		HandlerFunc: homeHandler,
	},
}

