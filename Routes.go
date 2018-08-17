package main

import (
	"net/http"
)

//Route : the route template.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes : all the route objects for Router.go
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"SightingIndex",
		"GET",
		"/sightings",
		SightingIndex,
	},
	Route{
		"SightingShow",
		"GET",
		"/sightings/{sightingId}",
		SightingShow,
	},
}
