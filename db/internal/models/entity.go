package models

import "time"

type University struct {
	UniversityID   int    `json:"university_id"`
	UniversityName string `json:"university_name"`
	EducationForm  string `json:"education_form"`
	Course         int    `json:"course"`
	GroupName      string `json:"group_name"`
	Specialty      string `json:"specialty"`
	Profile        string `json:"profile"`
}

type Passport struct {
	PassportID        int       `json:"passport_id"`
	PassportSeries    string    `json:"passport_series"`
	PassportIssueDate time.Time `json:"passport_issue_date"`
}

type Student struct {
	StudentID        int        `json:"student_id"`
	FullName         string     `json:"full_name"`
	SNILS            string     `json:"snils,omitempty"`
	BirthDate        time.Time  `json:"birth_date"`
	Email            string     `json:"email"`
	PhoneNumber      string     `json:"phone_number"`
	TelegramUsername string     `json:"telegram_username"`
	UniversityID     int        `json:"university_id"`
	PassportID       int        `json:"passport_id"`
	University       University `json:"university"`
	Passport         Passport   `json:"passport"`
}
