package service

import "github.com/nats-io/nats.go"

type Service struct {
	NatsClient *nats.Conn
}

func New(
	natsClient *nats.Conn,
) *Service {
	return &Service{
		NatsClient: natsClient,
	}
}
