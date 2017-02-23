package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


func NewRouter() *mux.Router {
	// create the router
	router := mux.NewRouter().StrictSlash(true)

	// Serve root from static
	router.Handle("/", Logger(http.FileServer(http.Dir("./static/")), "Index"))

	// Serve static from /static
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", 
		Logger(http.FileServer(http.Dir("./static/")), "Static")))

	// spawn API routes
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(APIMiddleware(handler))

	}
	return router
}
