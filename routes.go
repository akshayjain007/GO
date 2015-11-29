package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Login",
		"POST",
		"/data",
		Login,
	},
	Route{
		"KILL",
		"POST",
		"/kill",
		Kill,
	},
}
