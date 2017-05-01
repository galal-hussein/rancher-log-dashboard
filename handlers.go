package main

import (
	"html/template"
	"net/http"

	"github.com/galal-hussein/rancher-log-dashboard/container"
	. "github.com/galal-hussein/rancher-log-dashboard/influx"
	"github.com/gorilla/mux"
)

type templateData struct {
	C            container.Container
	Containers   container.Containers
	TemplateType string
}

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	var td templateData
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	t.Execute(w, &td)
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	var td templateData
	td.Containers = GetContainers("healthcheck")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	td.TemplateType = "service"
	t.Execute(w, &td)
}

func HealthcheckContainer(w http.ResponseWriter, r *http.Request) {
	var ok bool
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	vars := mux.Vars(r)
	contName := vars["contName"]
	var td templateData
	td.Containers = GetContainers("healthcheck")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	if td.C, ok = td.Containers[contName]; ok {
		td.TemplateType = "container"
		t.Execute(w, &td)
	} else {
		w.WriteHeader(404)
	}
}

func IPsec(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	var td templateData
	td.Containers = GetContainers("ipsec")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	td.TemplateType = "service"
	t.Execute(w, &td)
}

func IPsecContainer(w http.ResponseWriter, r *http.Request) {
	var ok bool
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	vars := mux.Vars(r)
	contName := vars["contName"]
	var td templateData
	td.Containers = GetContainers("ipsec")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	if td.C, ok = td.Containers[contName]; ok {
		td.TemplateType = "container"
		t.Execute(w, &td)
	} else {
		w.WriteHeader(404)
	}
}

func NetworkServices(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	var td templateData
	td.Containers = GetContainers("network")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	td.TemplateType = "service"
	t.Execute(w, &td)
}

func NetworkServicesContainer(w http.ResponseWriter, r *http.Request) {
	var ok bool
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	vars := mux.Vars(r)
	contName := vars["contName"]
	var td templateData
	td.Containers = GetContainers("network")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	if td.C, ok = td.Containers[contName]; ok {
		td.TemplateType = "container"
		t.Execute(w, &td)
	} else {
		w.WriteHeader(404)
	}
}

func Scheduler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	var td templateData
	td.Containers = GetContainers("scheduler")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	td.TemplateType = "service"
	t.Execute(w, &td)
}

func SchedulerContainer(w http.ResponseWriter, r *http.Request) {
	var ok bool
	t, _ := template.ParseFiles("public/dashboard.gtpl")
	vars := mux.Vars(r)
	contName := vars["contName"]
	var td templateData
	td.Containers = GetContainers("scheduler")
	if td.Containers == nil {
		w.WriteHeader(404)
	}
	if td.C, ok = td.Containers[contName]; ok {
		td.TemplateType = "container"
		t.Execute(w, &td)
	} else {
		w.WriteHeader(404)
	}
}
