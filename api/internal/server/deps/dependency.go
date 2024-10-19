package deps

import (
	"dpo-document-api/config"
)

func Process(config *config.ServiceConfig) *Providers {
	return &Providers{
		ServerHTTP: config.HTTP,
		Logger:     config.Logger.InitializeGlobally(),
		NatsClient: config.Nats.New(),
	}
}
