package service

import (
	"github.com/lukasjarosch/go-docx"
	natsContracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
)

func (s *Service) InsertDocumentDataIntoTemplate(document *natsContracts.Document) error {
	rm := natsContracts.ConvertDocumentToPlaceholderMap(document)

	doc, err := docx.Open("./template/tmp1.docx")
	doc2, err := docx.Open("./template/tmp2.docx")
	if err != nil {
		return err
	}

	doc.ReplaceAll(rm)
	doc2.ReplaceAll(rm)

	if err != nil {
		return err
	}

	err = doc.WriteToFile("./ans/replaced.docx")
	doc2.WriteToFile("./ans/replaced2.docx")

	if err != nil {
		return err
	}

	return nil
}
