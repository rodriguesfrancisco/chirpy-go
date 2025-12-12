package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const port = ":8080"
	mux := http.NewServeMux()
	handler := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))
	apiCfg := apiConfig{}
	apiCfg.fileserverHits.Store(0)
	mux.Handle("/app/", apiCfg.middlewareMetricsCounter(handler))
	mux.HandleFunc("/healthz", healthzHandler)
	mux.HandleFunc("/metrics/", apiCfg.metricsHandler)
	mux.HandleFunc("/reset/", apiCfg.resetHandler)

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

func (cfg *apiConfig) middlewareMetricsCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	hits := fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, hits)
}

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	cfg.fileserverHits.Swap(0)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
