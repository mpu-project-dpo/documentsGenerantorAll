package http

import (
	"dpo-document-api/internal/model"
	"encoding/json"
	"net/http"
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
	doc.FIO = doc1.Patronymic + " "+ doc1.Name + " " + doc1.Surename
	doc.Birthdate = doc1.Birthdate 
	doc.Course = doc1.Course
	doc.Group = doc1.Group
	doc.Phone = doc1.Phone
	doc.University = doc1.University
	doc.Profile = doc1.Profile
	doc.Email = doc1.Email
	doc.SnilsId = doc1.SnilsId
	doc.Specialty = doc1.Specialty
	doc.PassportIssueDate =  doc1.PassportIssueDate
	doc.PassportId = doc1.PassportId
	doc.CurrentAdress = doc1.CurrentAdress
	doc.RegAddress = doc1.RegAddress
	if err := h.DpoDocumentService.ProcessDocument(doc); err != nil {
		zap.L().Sugar().Errorf("Internal error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.DpoDocumentService.ProcessDocument(doc)
	w.WriteHeader(http.StatusOK)
}
