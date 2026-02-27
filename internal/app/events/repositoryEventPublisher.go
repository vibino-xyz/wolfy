package events

import (
	"context"

	repositoryv1 "github.com/vibino-xyz/protos/contracts/build/go/repository/v1"
)

type RepositoryEventPublisher interface {
	PublishRepositoryEvent(ctx context.Context, message *repositoryv1.RepositoryEventMessage) error
}
