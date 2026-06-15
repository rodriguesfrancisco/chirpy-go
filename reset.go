package main

import (
	"io"
	"net/http"
)

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Cannot run this operation", nil)
		return
	}
	cfg.databaseQueries.DeleteUsers(r.Context())
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	cfg.fileserverHits.Swap(0)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
