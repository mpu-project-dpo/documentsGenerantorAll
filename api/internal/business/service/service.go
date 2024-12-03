package service

import (
	nats "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats/producer"
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
