module github.com/vibino-xyz/wolfy

go 1.25.0

replace github.com/vibino-xyz/protos => ../protos

require (
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/vibino-xyz/protos v0.0.0
)

require google.golang.org/protobuf v1.36.11

require (
	github.com/labstack/echo/v5 v5.0.4 // indirect
	go.uber.org/dig v1.19.0 // indirect
	go.uber.org/fx v1.24.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)
