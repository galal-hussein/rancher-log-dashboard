package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var (
	debug = flag.Bool("debug", false, "Debug")
)

func getEnv() {
	logrus.Info("Starting Rancher infracture stack dashboard")
	flag.Parse()
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	getEnv()

	router := NewRouter()
	router.PathPrefix("/static/").Handler(http.FileServer(http.Dir("public/")))
	log.Fatal(http.ListenAndServe(":5000", router))
}
