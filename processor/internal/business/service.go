package business

import (
	natsContracts "nats-contracts"
)

type IDpoDocumentProcessingService interface {
	InsertDocumentDataIntoTemplate(document *natsContracts.Document) error
}
