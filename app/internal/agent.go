package agent

import (
	"fmt"
	"strings"

	"github.com/kakeru-lab/rpi-edge-agent/internal/memory"
	"github.com/kakeru-lab/rpi-edge-agent/internal/skills"
)

type Agent struct {
	store *memory.Store
	llm   *llmClient
}

func New(store *memory.Store) *Agent {
	return &Agent{
		store: store,
		llm:   newLLMClient(),
	}
}

func (a *Agent) Ask(sessionID, message string) (string, error) {
	// Save user message
	if err := a.store.AddMessage(sessionID, "user", message); err != nil {
		return "", err
	}

	system := `You are an edge AI assistant running on a Raspberry Pi.
Be concise. If tool results are provided, use them. Do not hallucinate device state.`

	// Tool routing (MVP): CPU temp only
	lower := strings.ToLower(message)
	toolInfo := ""
	if strings.Contains(lower, "cpu") || strings.Contains(lower, "temp") || strings.Contains(message, "温度") {
		t, err := skills.CPUTempCelsius()
		if err != nil {
			toolInfo = fmt.Sprintf("TOOL(cpu_temp) error: %v", err)
		} else {
			toolInfo = fmt.Sprintf("TOOL(cpu_temp) result: %.1f°C", t)
		}
	}

	user := message
	if toolInfo != "" {
		user = message + "\n\n" + toolInfo
	}

	reply, err := a.llm.Chat(system, user)
	if err != nil {
		// fallback response if LLM is not configured
		reply = "LLM is not configured. Set OPENAI_API_KEY / OPENAI_BASE_URL / OPENAI_MODEL."
	}

	if err := a.store.AddMessage(sessionID, "assistant", reply); err != nil {
		return "", err
	}
	return reply, nil
}
