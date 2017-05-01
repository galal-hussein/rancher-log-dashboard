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
		"Healthcheck Container",
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
		"IPsec Container",
		"GET",
		"/ipsec/{contName}",
		IPsecContainer,
	},
	Route{
		"Network Services",
		"GET",
		"/network",
		NetworkServices,
	},
	Route{
		"Network Services Container",
		"GET",
		"/network/{contName}",
		NetworkServicesContainer,
	},
	Route{
		"Scheduler",
		"GET",
		"/scheduler",
		Scheduler,
	},
	Route{
		"Scheduler Container",
		"GET",
		"/scheduler/{contName}",
		SchedulerContainer,
	},
}
