package nats

import (
	"go.uber.org/zap"
	natsContracts "nats-contracts"
)

func (h *Handler) ProcessDocument(doc *natsContracts.Document) error {
	if err := h.DpoDocumentService.InsertDocumentDataIntoTemplate(doc); err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		return err
	}

	return nil
}
