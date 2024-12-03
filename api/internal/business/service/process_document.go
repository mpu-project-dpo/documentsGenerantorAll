package service

import (
	"encoding/json"
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
)

func (s *Service) ProcessDocument(document *natsContracts.Document) error {
	bytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	return s.Nats.Publish(bytes)
}
