package deps

import (
	httpTransport "dpo-document-api/pkg/app-kit/http-transport"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Providers struct {
	ServerHttp *httpTransport.Config
	Logger     *zap.Logger
	NatsClient *nats.Conn
}
