package server

import (
	appKit "app-kit/server"
	processorService "dpo-document-processor/internal/business/service"
	processorTransport "dpo-document-processor/internal/business/transport/nats"
	"dpo-document-processor/internal/server/deps"
	natsContracts "nats-contracts"
	natsTransport "nats/consumer"
)

func GetController(deps *deps.Providers) appKit.IController {
	service := processorTransport.New(
		processorService.New(),
	)

	return natsTransport.New[natsContracts.Document](
		deps.NatsTransport,
		service.ProcessDocument,
	)
}
