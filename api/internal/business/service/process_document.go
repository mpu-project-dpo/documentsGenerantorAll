package service

import (
	"encoding/json"
	natsContracts "nats-contracts"
)

func (s *Service) ProcessDocument(document *natsContracts.Document) error {
	bytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	return s.Nats.Publish(bytes)
}
