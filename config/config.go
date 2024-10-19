package config

import (
	httpTransport "dpo-document-api/pkg/app-kit/http-transport"
	"dpo-document-api/pkg/app-kit/logger"
	natsClient "dpo-document-api/pkg/nats-client"
)

type ServiceConfig struct {
	Http   *httpTransport.Config
	Logger *logger.Config
	Nats   *natsClient.Config
}
