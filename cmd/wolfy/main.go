package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	webhooks "github.com/vibino-xyz/wolfy/internal/controller/http/github"
	"github.com/vibino-xyz/wolfy/internal/infra/rabbitmq"
	"go.uber.org/fx"
)

func NewEcho() *echo.Echo {
	return echo.New()
}

func RegisterRoutes(e *echo.Echo, handler *webhooks.Handler) {
	e.POST("/api/webhooks/github", handler.Handle)
}

func main() {
	fx.New(
		fx.Provide(
			rabbitmq.NewConn,
			rabbitmq.NewRepositoryEventPublisher,
			webhooks.NewHandler,
			NewEcho,
		),
		fx.Invoke(RegisterRoutes),
		fx.Invoke(func(lc fx.Lifecycle, e *echo.Echo) {
			server := &http.Server{
				Addr:    ":8000",
				Handler: e,
			}
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						slog.Info("Server running on :8000")
						if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
							slog.Error("Server stopped", "error", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return server.Shutdown(ctx)
				},
			})
		}),
	).Run()
}
