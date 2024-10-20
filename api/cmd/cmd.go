package cmd

import (
	"dpo-document-api/config"
	controller "dpo-document-api/internal/server"
	deps2 "dpo-document-api/internal/server/deps"
	"dpo-document-api/pkg/app-kit/server"
)

func Run() {
	server.WithApp(
		server.WithController[config.ServiceConfig, deps2.Providers](
			deps2.Process, controller.GetController,
		),
	)
}
