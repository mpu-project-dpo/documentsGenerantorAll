package cmd

import (
	"app-kit/server"
	"dpo-document-processor/config"
	controller "dpo-document-processor/internal/server"
	"dpo-document-processor/internal/server/deps"
)

func Run() {
	server.WithApp(
		server.WithController[config.ServiceConfig, deps.Providers](
			deps.Process, controller.GetController,
		),
	)
}
