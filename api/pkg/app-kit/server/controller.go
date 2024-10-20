package server

import (
	"dpo-document-api/pkg/app-kit/server/config"
)

func WithController[C any, D any](deps func(*C) *D, transport func(*D) IController) IController {
	return transport(deps(config.Parse[C]()))
}
