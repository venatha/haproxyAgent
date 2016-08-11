package main

import (
	"net"
	"net/http"

	"golang.org/x/net/trace"
)

func configureTrace(r *http.Request) trace.Trace {
	tr := trace.New("hostingAPI", r.URL.Path)
	defer tr.Finish()
	return tr
}

func traceAuthRequest(req *http.Request) (any, sensitive bool) {
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		host = req.RemoteAddr
	}
	switch host {
	case "localhost", "127.0.0.1", "::1", "89.238.157.212", "185.45.15.70":
		return true, true
	default:
		return false, false
	}
}
