package business

import (
	natsContracts "nats-contracts"
)

type IDpoDocumentService interface {
	ProcessDocument(document *natsContracts.Document) error
}
