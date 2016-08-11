package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func authHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("X-API-Key")

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			authResponder(w)
			return
		}

		ip := net.ParseIP(host)
		if ip == nil {
			authResponder(w)
			return
		}
		fmt.Println(ip)

		for _, cidr := range appConfig.AllowedHosts {
			_, net, err := net.ParseCIDR(cidr)
			if err != nil {
				continue
			}
			if net.Contains(ip) {
				h.ServeHTTP(w, r)
				return
			}
		}

		if auth == "" {
			authResponder(w)
			return
		}
		if auth != "abc123" {
			authResponder(w)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func authResponder(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid API Key Passed"))
}

func genericResponse(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-This-Host", appConfig.ThisHostname)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func genericLogger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)
		appLog.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
	return http.HandlerFunc(fn)
}
