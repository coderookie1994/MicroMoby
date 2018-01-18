package common

import "net/http"

type (

	// Route :
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	// Routes :
	Routes []Route
)
