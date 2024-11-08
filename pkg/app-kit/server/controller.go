package server

import (
	"github.com/mpu-project-dpo/documentsGenerantorAll/pkg/app-kit/server/config"
)

func WithController[C any, D any](deps func(*C) *D, transport func(*D) IController) IController {
	return transport(deps(config.Parse[C]()))
}
