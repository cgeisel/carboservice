package main

import "net/http"

// Route has Name, Http Method, Pattern (Path), HandlerFunc for processing request
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of Route structs
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CarbsIndex",
		"GET",
		"/carbs",
		CarbsIndex,
	},
	Route{
		"CarbsOne",
		"GET",
		"/carbs/{ndbID}",
		CarbsOne,
	},
}
