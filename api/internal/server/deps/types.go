package deps

import (
	httpTransport "app-kit/http-transport"
	"go.uber.org/zap"
	natsClient "nats/producer"
)

type Providers struct {
	ServerHTTP *httpTransport.Config
	Logger     *zap.Logger
	NatsClient *natsClient.Client
}
