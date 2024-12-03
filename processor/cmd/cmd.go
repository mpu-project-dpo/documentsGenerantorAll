package cmd

import (
	"dpo-document-processor/config"
	controller "dpo-document-processor/internal/server"
	"dpo-document-processor/internal/server/deps"
	"github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/server"
)

func Run() {
	server.WithApp(
		server.WithController[config.ServiceConfig, deps.Providers](
			deps.Process, controller.GetController,
		),
	)
}
