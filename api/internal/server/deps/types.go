package deps

import (
	httpTransport "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/http-transport"
	natsClient "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats/producer"
	"go.uber.org/zap"
)

type Providers struct {
	ServerHTTP *httpTransport.Config
	Logger     *zap.Logger
	NatsClient *natsClient.Client
}
