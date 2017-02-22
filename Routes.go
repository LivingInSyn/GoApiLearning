package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route identifies a full route to add to the router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of the Route struct
type Routes []Route

// NewRouter returns a new gorilla mux router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"EventIndex",
		"GET",
		"/events",
		EventIndex,
	},
	Route{
		"EventShow",
		"GET",
		"/events/{eventId}",
		EventShow,
	},
	Route{
		"EventAdd",
		"POST",
		"/events/add",
		EventAdd,
	},
}
