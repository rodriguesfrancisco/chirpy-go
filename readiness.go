package main

import (
	"io"
	"net/http"
)

func healthzHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "OK")
}
