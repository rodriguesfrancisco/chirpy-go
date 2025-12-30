package main

import (
	"fmt"
	"io"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	hitsTemplate := `<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>`
	hits := fmt.Sprintf(hitsTemplate, cfg.fileserverHits.Load())
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, hits)
}
