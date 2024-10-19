package cmd

import (
	"dpo-document-api/config"
	controller "dpo-document-api/internal/server"
	"dpo-document-api/internal/server/deps"
	appKit "dpo-document-api/pkg/app-kit/server"
)

func Run() {
	appKit.WithApp(
		appKit.WithController[config.ServiceConfig, deps.Providers](
			deps.Process, controller.GetController,
		),
	)
}
