package config

import (
	httpTransport "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/http-transport"
	"github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/logger"
	natsClient "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats/producer"
)

type ServiceConfig struct {
	HTTP   *httpTransport.Config
	Logger *logger.Config
	Nats   *natsClient.Config
}
