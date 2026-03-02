package github

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/vibino-xyz/wolfy/internal/app/events"
)

type GithubHandler struct {
	repositoryEventPublisher events.RepositoryEventPublisher
}

func NewGithubHandler(repositoryEventPublisher events.RepositoryEventPublisher) *GithubHandler {
	return &GithubHandler{
		repositoryEventPublisher: repositoryEventPublisher,
	}
}

func (h *GithubHandler) Handle(c *echo.Context) error {
	slog.Info("Received github webhook")

	var payload GitHubWebhookPayload
	if err := c.Bind(&payload); err != nil {
		slog.Error("Failed to bind JSON payload", "error", err)
		return c.NoContent(http.StatusBadRequest)
	}

	protoMessage := ToProtos(payload)
	if err := h.repositoryEventPublisher.PublishRepositoryEvent(c.Request().Context(), protoMessage); err != nil {
		slog.Error("Failed to publish repository event", "error", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	slog.Info("Received GitHub webhook", "repository", payload.Repository.FullName, "hook_type", payload.Hook.Type)
	return c.NoContent(http.StatusOK)
}
