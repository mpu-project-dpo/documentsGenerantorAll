package config

import (
	"github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/logger"
	natsTransport "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats/consumer"
)

type ServiceConfig struct {
	NatsTransport *natsTransport.Config
	Logger        *logger.Config
}
