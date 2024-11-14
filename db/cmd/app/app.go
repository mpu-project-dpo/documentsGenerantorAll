package app

import (
	"ape/internal/db"
	"ape/internal/models"
	"ape/internal/nats"
	"log"

	"github.com/spf13/viper"
)

func Run() {
	// Инициализация Viper для чтения конфигурационного файла
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Создаем подключение к базе данных
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	// Создаем буферизованный канал для передачи студентов
	studentChan := make(chan models.Student, 10)

	// Запускаем горутину для прослушивания сообщений из NATS
	go func() {
		if err := nats.ListenAndProcessMessages(studentChan); err != nil {
			log.Fatalf("Failed to listen to NATS messages: %v", err)
		}
	}()

	// Запускаем горутину для обработки и сохранения студентов в базе данных
	go conn.ProcessAndSaveStudents(studentChan)

	// Блокируем основной поток, чтобы программа не завершилась
	select {}
}
