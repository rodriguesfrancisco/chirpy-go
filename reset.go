package main

import (
	"io"
	"net/http"
)

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	cfg.fileserverHits.Swap(0)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
