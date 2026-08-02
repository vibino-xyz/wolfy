package github

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/vibino-xyz/wolfy/internal/app/events"
)

type Handler struct {
	repositoryEventPublisher events.RepositoryEventPublisher
}

func NewHandler(repositoryEventPublisher events.RepositoryEventPublisher) *Handler {
	return &Handler{
		repositoryEventPublisher: repositoryEventPublisher,
	}
}

func (h *Handler) Handle(c *echo.Context) error {
	slog.Info("Received github webhook")

	var payload PushPayload
	if err := c.Bind(&payload); err != nil {
		slog.Error("Failed to bind JSON payload", "error", err)
		return c.NoContent(http.StatusBadRequest)
	}

	protoMessage := ToProto(payload)
	if err := h.repositoryEventPublisher.PublishRepositoryEvent(c.Request().Context(), protoMessage); err != nil {
		slog.Error("Failed to publish repository event", "error", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	slog.Info("Received GitHub webhook", "repository", payload.Repository.FullName)
	return c.NoContent(http.StatusOK)
}
