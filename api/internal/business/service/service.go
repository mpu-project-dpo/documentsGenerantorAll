package service

import (
	nats "nats/producer"
)

type Service struct {
	Nats *nats.Client
}

func New(
	nats *nats.Client,
) *Service {
	return &Service{
		Nats: nats,
	}
}
