package http

import (
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) ProcessDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		zap.L().Error("file id is empty")
		return
	}

	err := h.DpoDocumentService.ProcessDocument(doc)
	if err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}
