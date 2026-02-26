package webhooks

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type GithubHandler struct {
}

func NewGithubHandler() *GithubHandler {
	return &GithubHandler{}
}

func (h *GithubHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Failed to read request body", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))

	var payload GitHubWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		slog.Error("Failed to unmarshal JSON payload", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Info("Received GitHub webhook", "repository", payload.Repository.FullName, "hook_type", payload.Hook.Type)

	w.WriteHeader(http.StatusOK)
}
