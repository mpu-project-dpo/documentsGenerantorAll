package http

import (
	"dpo-document-api/internal/business"
)

type Handler struct {
	DpoDocumentService business.IDpoDocumentService
}

func New(
	s business.IDpoDocumentService,
) *Handler {
	return &Handler{DpoDocumentService: s}
}
