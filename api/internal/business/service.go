package business

import (
	"dpo-document-api/internal/business/models"
)

type IDpoDocumentService interface {
	ProcessDocument(document *models.Document) error
}
