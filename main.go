package main

import (
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
	mux.HandleFunc("GET /api/healthz", healthzHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetHandler)

	srv := http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
