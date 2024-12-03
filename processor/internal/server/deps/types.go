package deps

import (
	natsTransport "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats/consumer"
	"go.uber.org/zap"
)

type Providers struct {
	NatsTransport *natsTransport.Config
	Logger        *zap.Logger
}
