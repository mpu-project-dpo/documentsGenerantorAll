package deps

import (
	"go.uber.org/zap"
	natsTransport "nats/consumer"
)

type Providers struct {
	NatsTransport *natsTransport.Config
	Logger        *zap.Logger
}
