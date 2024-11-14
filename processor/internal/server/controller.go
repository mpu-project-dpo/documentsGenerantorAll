package server

import (
	processorService "dpo-document-processor/internal/business/service"
	processorTransport "dpo-document-processor/internal/business/transport/nats"
	"dpo-document-processor/internal/server/deps"
	appKit "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/server"
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
	natsTransport "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats/consumer"
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
