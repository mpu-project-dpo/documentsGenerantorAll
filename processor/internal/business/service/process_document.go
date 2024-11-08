package service

import (
	"github.com/lukasjarosch/go-docx"
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
)

func (s *Service) InsertDocumentDataIntoTemplate(document *natsContracts.Document) error {
	rm := natsContracts.ConvertDocumentToPlaceholderMap(document)

	doc, err := docx.Open("./template/template.docx")
	if err != nil {
		return err
	}

	err = doc.ReplaceAll(rm)
	if err != nil {
		return err
	}

	err = doc.WriteToFile("replaced.docx")
	if err != nil {
		return err
	}

	return nil
}
