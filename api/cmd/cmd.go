package cmd

import (
	"app-kit/server"
	"dpo-document-api/config"
	controller "dpo-document-api/internal/server"
	"dpo-document-api/internal/server/deps"
)

func Run() {
	server.WithApp(
		server.WithController[config.ServiceConfig, deps.Providers](
			deps.Process, controller.GetController,
		),
	)
}
