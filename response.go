package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/trace"
)

// HTTPResponseWriter returns a response from the API
func HTTPResponseWriter(w http.ResponseWriter, tr trace.Trace, err error, responseCode int, message string) {
	var response HTTPResponse
	tr.LazyPrintf(message+": %v", err)
	if err != nil {
		tr.SetError()
	}
	w.WriteHeader(responseCode)
	response = HTTPResponse{
		ResponseCode: responseCode,
		Response:     message}
	json.NewEncoder(w).Encode(response)
	return
}
