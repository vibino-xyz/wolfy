package rabbitmq

import (
	"context"
	"fmt"

	"github.com/Vibino/wolfy/internal/core"
	amqp "github.com/rabbitmq/amqp091-go"
)

type repositoryEventPublisher struct {
	ch       *amqp.Channel
	exchange string
}

const (
	QueueName    = "repository_event_queue"
	MainExchange = "repository_event_exchange"
	DLXExchange  = "repository_event_dlx_exchange"
)

func NewRepositoryEventPublisher(conn *amqp.Connection) (core.RepositoryEventPublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		MainExchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare main exchange: %w", err)
	}

	err = ch.ExchangeDeclare(
		DLXExchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare DLX exchange: %w", err)
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    DLXExchange,
		"x-dead-letter-routing-key": "failed",
	}
	queue, err := ch.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		queue.Name,
		"",
		MainExchange,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue with main exchange: %w", err)
	}

	return &repositoryEventPublisher{
		ch:       ch,
		exchange: MainExchange,
	}, nil
}

func (r *repositoryEventPublisher) PublishRepositoryEvent(ctx context.Context, req *core.RepositoryEventRequest) error {
	panic("unimplemented")
}
