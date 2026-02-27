module github.com/vibino-xyz/wolfy

go 1.24.0

replace github.com/vibino-xyz/protos => ../protos

require (
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/vibino-xyz/protos v0.0.0
)

require google.golang.org/protobuf v1.36.11
