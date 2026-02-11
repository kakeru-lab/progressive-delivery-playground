package api

import "net/http"

func Router(h *Handlers) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.Healthz)
	mux.HandleFunc("/agent/ask", h.Ask)
	return mux
}
