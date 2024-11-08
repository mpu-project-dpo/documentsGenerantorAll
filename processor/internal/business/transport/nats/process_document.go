package nats

import (
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
	"go.uber.org/zap"
)

func (h *Handler) ProcessDocument(doc *natsContracts.Document) error {
	if err := h.DpoDocumentService.InsertDocumentDataIntoTemplate(doc); err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		return err
	}

	return nil
}
