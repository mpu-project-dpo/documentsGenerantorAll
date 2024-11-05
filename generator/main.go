package main

import (
	"log"

	"baliance.com/gooxml/document"
)

func main() {
	doc := document.New()

	addFormattedParagraph(doc, "Ректору", 12, false, false)
	addFormattedParagraph(doc, "федерального государственного автономного образовательного учреждения", 12, false, false)
	addFormattedParagraph(doc, "высшего образования", 12, false, false)
	addFormattedParagraph(doc, "«Московский политехнический университет»", 12, false, false)
	addFormattedParagraph(doc, "В.В. Миклушевскому", 12, false, false)

	addFormattedParagraph(doc, "\tот _____________________________________", 12, false, false)
	addFormattedParagraph(doc, "\tФамилия Имя Отчество", 12, false, false)
	addFormattedParagraph(doc, "\t____________________________________ телефон, e-mail", 12, false, false)

	addFormattedParagraph(doc, "ЗАЯВЛЕНИЕ", 16, true, false)

	addFormattedParagraph(doc, "Прошу зачислить меня на дополнительную профессиональную программу профессиональной переподготовки", 12, false, false)
	addFormattedParagraph(doc, "«___________________________________________________________________»", 12, false, false)

	addFormattedParagraph(doc, "\tО себе сообщаю следующее:", 12, true, false)
	addFormattedParagraph(doc, "1. Ф. И. О. (полностью)________________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "2. Дата рождения______________________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "3. Паспортные данные_________________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "(номер, серия, кем выдан, дата выдачи, код подразделения)", 12, false, true)

	addFormattedParagraph(doc, "4. Адрес регистрации __________________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "Адрес проживания____________________________________________________________________", 12, false, false)

	addFormattedParagraph(doc, "5. Сведения об образовании_____________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "6. СНИЛС: ___________________________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "7. ИНН: _____________________________________________________________________________", 12, false, false)
	addFormattedParagraph(doc, "8. Занимаемая должность на момент обучения студент Московского Политехнического Университета, факультет/институт_______________________________________________", 12, false, false)
	addFormattedParagraph(doc, "Направление подготовки/специальность___________________________________________", 12, false, false)
	addFormattedParagraph(doc, "Профиль подготовки_______________________________ Группа______________________", 12, false, false)

	addFormattedParagraph(doc, "_________________________ .20____ г.", 12, false, false)
	addFormattedParagraph(doc, "подпись дата", 12, false, false)

	err := doc.SaveToFile("output.docx")
	if err != nil {
		log.Fatalf("Ошибка сохранения файла: %v", err)
	}
	log.Println("Файл успешно создан: output.docx")
}

func addFormattedParagraph(doc *document.Document, text string, fontSize int, bold bool, italic bool) {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(text)
	run.Properties().SetFontFamily("Times New Roman")
	run.Properties().SetBold(bold)
	run.Properties().SetItalic(italic)
}
