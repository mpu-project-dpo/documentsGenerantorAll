package main

import (
	"ape/internal/db"
	"ape/internal/nats"
	"log"
)

func main() {
	// Подключаемся к базе данных
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close() // Закрываем пул соединений при завершении работы

	// Запускаем слушатель NATS JetStream
	err = nats.ListenAndProcessMessages(conn)
	if err != nil {
		log.Fatalf("Failed to listen to NATS messages: %v", err)
	}
}
