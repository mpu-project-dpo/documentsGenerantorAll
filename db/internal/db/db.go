package db

import (
	"ape/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type DBConnection struct {
	Conn *pgx.Conn
}

func ConnectDB() (DBConnection, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/mydb")
	if err != nil {
		return DBConnection{}, err
	}
	return DBConnection{Conn: conn}, nil
}

// Метод для закрытия соединения
func (db DBConnection) Close() {
	err := db.Conn.Close(context.Background())
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

// Сохранение данных студента в базе данных
func (db DBConnection) SaveStudent(student models.Student) error {
	query := `
    INSERT INTO students (
        full_name, university, education_form, course, "group", specialty, profile,
        snils, passport_series, passport_issue_date, birth_date, email, phone_number, telegram_username
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := db.Conn.Exec(context.Background(), query,
		student.FullName, student.University, student.EducationForm, student.Course, student.Group,
		student.Specialty, student.Profile, student.SNILS, student.PassportSeries, student.PassportIssueDate,
		student.BirthDate, student.Email, student.PhoneNumber, student.TelegramUsername)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	return nil
}
