package business

import (
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
)

type IDpoDocumentService interface {
	ProcessDocument(document *natsContracts.Document) error
}
