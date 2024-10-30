package db

import (
	"ape/internal/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type DBConnection struct {
	Pool *pgxpool.Pool
}

// Подключение к базе данных PostgreSQL с использованием настроек из конфигурационного файла
func ConnectDB() (DBConnection, error) {
	// Инициализируем Viper и загружаем конфигурацию
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Получаем URL базы данных из конфигурации
	dbURL := viper.GetString("database.url")

	// Подключаемся к базе данных с использованием пула соединений
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return DBConnection{}, err
	}

	return DBConnection{Pool: pool}, nil
}

// Метод для закрытия пула соединений
func (db DBConnection) Close() {
	db.Pool.Close()
}

// Сохранение данных студента в базе данных
func (db DBConnection) SaveStudent(ctx context.Context, student models.Student) error {
	query := `
    INSERT INTO students (
        full_name, university, education_form, course, "group", specialty, profile,
        snils, passport_series, passport_issue_date, birth_date, email, phone_number, telegram_username
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := db.Pool.Exec(ctx, query,
		student.FullName, student.University, student.EducationForm, student.Course, student.Group,
		student.Specialty, student.Profile, student.SNILS, student.PassportSeries, student.PassportIssueDate,
		student.BirthDate, student.Email, student.PhoneNumber, student.TelegramUsername)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	return nil
}
