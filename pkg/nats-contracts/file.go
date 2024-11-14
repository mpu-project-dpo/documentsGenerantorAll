package nats_contracts

import "time"

type Document struct {
	FIO               string    `json:"fio,omitempty"`
	University        string    `json:"university,omitempty"`
	StudyForm         string    `json:"study_form,omitempty"`
	Course            string    `json:"course,omitempty"`
	Group             string    `json:"group,omitempty"`
	Specialty         string    `json:"specialty,omitempty"`
	Profile           string    `json:"profile,omitempty"`
	SnilsId           string    `json:"snils_id,omitempty"`
	PassportId        string    `json:"passport_id,omitempty"`
	PassportIssueDate time.Time `json:"passport_issue_date,omitempty"`
	Birthdate         time.Time `json:"birthdate,omitempty"`
	Email             string    `json:"email,omitempty"`
	Phone             string    `json:"phone,omitempty"`
	TelegramId        string    `json:"telegram_id"`
}
