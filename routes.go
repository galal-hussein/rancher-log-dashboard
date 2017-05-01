package main

import (
	"net/http"
)

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
		"Healthcheck",
		"GET",
		"/healthcheck",
		Healthcheck,
	},
	Route{
		"HealthcheckContainer",
		"GET",
		"/healthcheck/{contName}",
		HealthcheckContainer,
	},
	Route{
		"IPsec",
		"GET",
		"/ipsec",
		IPsec,
	},
	Route{
		"Network Services",
		"GET",
		"/network",
		NetworkServices,
	},
	Route{
		"Scheduler",
		"GET",
		"/scheduler",
		Scheduler,
	},
}
