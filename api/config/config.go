package config

import (
	httpTransport "app-kit/http-transport"
	"app-kit/logger"
	natsClient "nats/producer"
)

type ServiceConfig struct {
	HTTP   *httpTransport.Config
	Logger *logger.Config
	Nats   *natsClient.Config
}
