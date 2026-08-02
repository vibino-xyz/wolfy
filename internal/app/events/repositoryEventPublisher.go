package events

import (
	"context"

	contracts "github.com/vibino-xyz/protos/contracts/build"
)

type RepositoryEventPublisher interface {
	PublishRepositoryEvent(ctx context.Context, message *contracts.RepositoryEventMessage) error
}
