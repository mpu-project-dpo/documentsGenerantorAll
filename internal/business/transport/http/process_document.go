package http

import (
	"dpo-document-api/internal/business/models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) ProcessDocument(w http.ResponseWriter, r *http.Request) {
	doc := new(models.Document)

	if err := json.NewDecoder(r.Body).Decode(doc); err != nil {
		return
	}

	if err := h.DpoDocumentService.ProcessDocument(doc); err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}
