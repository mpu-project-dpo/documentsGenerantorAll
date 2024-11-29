package http

import (
	"dpo-document-api/internal/model"
	"encoding/json"
	"net/http"
	"time"

	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
	"go.uber.org/zap"
)

func (h *Handler) ProcessDocument(w http.ResponseWriter, r *http.Request) {
	doc1 := new(model.Document)

	if err := json.NewDecoder(r.Body).Decode(doc1); err != nil {
		zap.L().Sugar().Errorf("decode error %s", err.Error())
		return
	}
	doc := new(natsContracts.Document)
	doc.FIO = doc1.Name + " " + doc1.Surename
	doc.Birthdate, _ = time.Parse("2/1/2006", doc1.Birthdate) 
	doc.Course = doc1.Course
	doc.Group = doc1.Group
	doc.Phone = doc1.Phone
	doc.University = doc1.University
	doc.Profile = doc1.Profile
	doc.Email = doc1.Email
	doc.SnilsId = doc1.SnilsId
	doc.Specialty = doc1.Specialty
	doc.PassportIssueDate, _ = time.Parse("2/1/2006", doc1.PassportIssueDate)
	if err := h.DpoDocumentService.ProcessDocument(doc); err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.DpoDocumentService.ProcessDocument(doc)
	w.WriteHeader(http.StatusOK)
}
