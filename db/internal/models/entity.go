package models

import "time"

type Student struct {
	FullName          string    `json:"full_name"`
	University        string    `json:"university"`
	EducationForm     string    `json:"education_form"`
	Course            int       `json:"course"`
	Group             string    `json:"group"`
	Specialty         string    `json:"specialty"`
	Profile           string    `json:"profile"`
	SNILS             string    `json:"snils,omitempty"`
	PassportSeries    string    `json:"passport_series"`
	PassportIssueDate time.Time `json:"passport_issue_date"`
	BirthDate         time.Time `json:"birth_date"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phone_number"`
	TelegramUsername  string    `json:"telegram_username"`
}
