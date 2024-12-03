package model


type Document struct {
	Name               string    `json:"name,omitempty"` 
	Surename               string    `json:"surname,omitempty"` 
	University        string    `json:"university,omitempty"` 
	StudyForm         string    `json:"study_form,omitempty"` 
	Course            string    `json:"course,omitempty"` 
	Group             string    `json:"group,omitempty"` 
	Specialty         string    `json:"specialty,omitempty"` 
	Profile           string    `json:"dpo,omitempty"` 
	SnilsId           string    `json:"snils,omitempty"` 
	PassportId        string    `json:"passport,omitempty"` 
	Patronymic string `json:"patronymic,omitempty"`
	PassportIssueDate string `json:"issuedBy,omitempty"` 
	Birthdate         string `json:"dateOfBirth,omitempty"` 
	Email             string    `json:"email,omitempty"` 
	Phone             string    `json:"phone,omitempty"` 
	TelegramId        string    `json:"telegram_id,omitempty"`
	CurrentAdress string `json:"currentAddress,omitempty"`
	RegAddress string `json:"regAddress,omitempty"`
}
