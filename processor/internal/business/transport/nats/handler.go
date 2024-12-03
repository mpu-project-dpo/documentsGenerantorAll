package nats

import (
	"dpo-document-processor/internal/business"
)

type Handler struct {
	DpoDocumentService business.IDpoDocumentProcessingService
}

func New(
	s business.IDpoDocumentProcessingService,
) *Handler {
	return &Handler{DpoDocumentService: s}
}
