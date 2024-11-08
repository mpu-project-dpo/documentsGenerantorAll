package config

import (
	"app-kit/logger"
	natsTransport "nats/consumer"
)

type ServiceConfig struct {
	NatsTransport *natsTransport.Config
	Logger        *logger.Config
}
