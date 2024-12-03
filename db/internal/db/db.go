package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	nats_contracts "github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
	"github.com/spf13/viper"
)

// ProcessAndSaveStudents обрабатывает студентов из канала и сохраняет их в базе данных
func (db DBConnection) ProcessAndSaveStudents(c <-chan nats_contracts.Document) {
	for document := range c {
		if err := db.SaveDocument(document); err != nil {
			log.Printf("Error saving document: %v", err)
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
func (db DBConnection) SaveUniversity(tx pgx.Tx, university string, studyForm string, course string, group string, specialty string, profile string) (int, error) {
	var universityID int
	err := tx.QueryRow(context.Background(),
		"INSERT INTO universities (university_name, education_form, course, group_name, specialty, profile) VALUES ($1, $2, $3, $4, $5, $6) RETURNING university_id",
		university, studyForm, course, group, specialty, profile,
	).Scan(&universityID)
	return universityID, err
}

// SavePassport сохраняет данные паспорта в базу данных
func (db DBConnection) SavePassport(tx pgx.Tx, passportId string, passportIssueDate time.Time) (int, error) {
	var passportID int
	err := tx.QueryRow(context.Background(),
		"INSERT INTO passports (passport_series, passport_issue_date) VALUES ($1, $2) RETURNING passport_id",
		passportId, passportIssueDate,
	).Scan(&passportID)
	return passportID, err
}

// SaveDocument сохраняет данные документа в базу данных
func (db DBConnection) SaveDocument(document nats_contracts.Document) error {
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Сохраняем данные университета
	universityID, err := db.SaveUniversity(tx, document.University, document.StudyForm, document.Course, document.Group, document.Specialty, document.Profile)
	if err != nil {
		return err
	}

	// Сохраняем данные паспорта
	passportID, err := db.SavePassport(tx, document.PassportId, document.PassportIssueDate)
	if err != nil {
		return err
	}

	// Сохраняем данные студента
	_, err = tx.Exec(context.Background(),
		"INSERT INTO students (full_name, snils, birth_date, email, phone_number, telegram_username, university_id, passport_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		document.FIO, document.SnilsId, document.Birthdate, document.Email, document.Phone, document.TelegramId, universityID, passportID,
	)
	if err != nil {
		return err
	}

	// Фиксируем транзакцию
	return tx.Commit(context.Background())
}
