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

// SaveUniversity сохраняет данные университета в базу данных
func (db DBConnection) SaveUniversity(university models.University) (int, error) {
	var universityID int
	err := db.Pool.QueryRow(context.Background(),
		"INSERT INTO universities (university_name, education_form, course, group_name, specialty, profile) VALUES ($1, $2, $3, $4, $5, $6) RETURNING university_id",
		university.UniversityName, university.EducationForm, university.Course, university.GroupName, university.Specialty, university.Profile,
	).Scan(&universityID)
	return universityID, err
}

// SavePassport сохраняет данные паспорта в базу данных
func (db DBConnection) SavePassport(passport models.Passport) (int, error) {
	var passportID int
	err := db.Pool.QueryRow(context.Background(),
		"INSERT INTO passports (passport_series, passport_issue_date) VALUES ($1, $2) RETURNING passport_id",
		passport.PassportSeries, passport.PassportIssueDate,
	).Scan(&passportID)
	return passportID, err
}

// SaveStudent сохраняет данные студента в базу данных
func (db DBConnection) SaveStudent(student models.Student) error {
	// Сохраняем данные университета
	universityID, err := db.SaveUniversity(student.University)
	if err != nil {
		return err
	}

	// Сохраняем данные паспорта
	passportID, err := db.SavePassport(student.Passport)
	if err != nil {
		return err
	}

	// Сохраняем данные студента
	_, err = db.Pool.Exec(context.Background(),
		"INSERT INTO students (full_name, snils, birth_date, email, phone_number, telegram_username, university_id, passport_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		student.FullName, student.SNILS, student.BirthDate, student.Email, student.PhoneNumber, student.TelegramUsername, universityID, passportID,
	)
	return err
}
