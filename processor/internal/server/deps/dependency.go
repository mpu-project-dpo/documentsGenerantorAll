package deps

import (
	"dpo-document-processor/config"
)

func Process(config *config.ServiceConfig) *Providers {
	return &Providers{
		NatsTransport: config.NatsTransport,
		Logger:        config.Logger.InitializeGlobally(),
	}
}
