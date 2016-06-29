package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type routes []route

var routeList = routes{
	route{
		"Index",
		"GET",
		"/",
		home,
	},
	route{
		"Game",
		"GET",
		"/game",
		battleship,
	},
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(staticHandler)

	for _, r := range routeList {
		router.
			Methods(r.method).
			Path(r.pattern).
			Name(r.name).
			Handler(r.handlerFunc)
	}

	return router
}
