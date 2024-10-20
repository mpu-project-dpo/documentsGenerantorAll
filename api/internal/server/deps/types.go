package deps

import (
	httpTransport "dpo-document-api/pkg/app-kit/http-transport"
	natsClient "dpo-document-api/pkg/nats-client"
	"go.uber.org/zap"
)

type Providers struct {
	ServerHTTP *httpTransport.Config
	Logger     *zap.Logger
	NatsClient *natsClient.Client
}
