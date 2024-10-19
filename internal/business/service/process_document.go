package service

import (
	"dpo-document-api/internal/business/models"
	"encoding/json"
)

func (s *Service) ProcessDocument(document *models.Document) error {
	bytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	return s.Nats.Publish(bytes)
}
