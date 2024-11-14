package business

import (
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
)

type IDpoDocumentProcessingService interface {
	InsertDocumentDataIntoTemplate(document *natsContracts.Document) error
}
