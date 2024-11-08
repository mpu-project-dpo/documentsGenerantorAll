package cmd

import (
	"dpo-document-api/config"
	controller "dpo-document-api/internal/server"
	"dpo-document-api/internal/server/deps"
	"github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/server"
)

func Run() {
	server.WithApp(
		server.WithController[config.ServiceConfig, deps.Providers](
			deps.Process, controller.GetController,
		),
	)
}
