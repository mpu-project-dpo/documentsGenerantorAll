package http

import (
	"encoding/json"
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) ProcessDocument(w http.ResponseWriter, r *http.Request) {
	doc := new(natsContracts.Document)

	if err := json.NewDecoder(r.Body).Decode(doc); err != nil {
		zap.L().Sugar().Errorf("decode error %s", err.Error())
		return
	}

	if err := h.DpoDocumentService.ProcessDocument(doc); err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.DpoDocumentService.ProcessDocument(doc)
	w.WriteHeader(http.StatusOK)
}
