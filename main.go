package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	const port = ":8080"
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", healthzHandler)

	srv := http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

func healthzHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "OK")
}
