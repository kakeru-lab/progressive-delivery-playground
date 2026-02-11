package api

import (
	"encoding/json"
	"net/http"

	"github.com/kakeru-lab/rpi-edge-agent/internal/agent"
)

type Handlers struct {
	agent *agent.Agent
}

func NewHandlers(a *agent.Agent) *Handlers {
	return &Handlers{agent: a}
}

func (h *Handlers) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

type askRequest struct {
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}

type askResponse struct {
	Reply string `json:"reply"`
}

func (h *Handlers) Ask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req askRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.SessionID == "" {
		req.SessionID = "default"
	}

	reply, err := h.agent.Ask(req.SessionID, req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(askResponse{Reply: reply})
}
