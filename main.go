// Copyright 2016 Dan Longshaw. All rights reserved.
/*
	This application serves as a REST API for a custom built haproxy management platform
	It will implement basic GET/POST/DELETE routes for most haproxy functions
*/
package main

import (
	"log"
	"net/http"

	"golang.org/x/net/trace"

	"goji.io"
	"goji.io/pat"
)

var appConfig *AppConfig
var appLog *log.Logger

func init() {
	trace.AuthRequest = traceAuthRequest
}

func main() {
	appConfig = loadConfig()
	appLog = loadAppLogger()

	mux := goji.NewMux()
	mux.Use(authHandler)
	mux.Use(genericResponse)
	mux.Use(genericLogger)

	mux.HandleFunc(pat.Get("/index"), indexHandler)
	http.Handle("/", mux)

	appLog.Println("Starting HTTP server")
	appLog.Fatal(http.ListenAndServe(":3000", nil))
}
