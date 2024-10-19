package service

import (
	nats "dpo-document-api/pkg/nats-client"
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
