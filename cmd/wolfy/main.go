package main

import (
	"log"
	"log/slog"
	"net/http"

	webhooks "github.com/vibino-xyz/wolfy/internal/controller/http/github"
)

func main() {
	handler := webhooks.NewGithubHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/webhooks/github", handler.Handle)

	slog.Info("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Failed to start server", "error", err)
		log.Fatal(err)
	}
}
