package db

import (
	"ape/internal/models"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

// ProcessAndSaveStudents обрабатывает студентов из канала и сохраняет их в базе данных
func (db DBConnection) ProcessAndSaveStudents(c <-chan models.Student) {
	for student := range c {
		if err := db.SaveStudent(student); err != nil {
			log.Printf("Error saving student: %v", err)
		}
	}
}

type DBConnection struct {
	Pool *pgxpool.Pool
}

// ConnectDB создает соединение с базой данных с использованием пула соединений
func ConnectDB() (DBConnection, error) {
	dbURL := viper.GetString("database.url")
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return DBConnection{}, err
	}

	// Устанавливаем тайм-ауты и параметры пула
	config.MaxConns = 10
	config.HealthCheckPeriod = time.Minute

	// Подключаемся к базе данных с использованием пула соединений
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return DBConnection{}, err
	}

	return DBConnection{Pool: pool}, nil
}

// Close закрывает пул соединений
func (db DBConnection) Close() {
	db.Pool.Close()
}

// SaveStudent сохраняет данные студента в базу данных
func (db DBConnection) SaveStudent(student models.Student) error {
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO students (full_name, university, education_form, course, group, specialty, profile, snils, passport_series, passport_issue_date, birth_date, email, phone_number, telegram_username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		student.FullName, student.University, student.EducationForm, student.Course, student.Group, student.Specialty, student.Profile, student.SNILS, student.PassportSeries, student.PassportIssueDate, student.BirthDate, student.Email, student.PhoneNumber, student.TelegramUsername,
	)
	return err
}
