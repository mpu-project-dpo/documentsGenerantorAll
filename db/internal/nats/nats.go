package nats

import (
	"ape/internal/db"
	"ape/internal/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const natsSubject = "students_data"

// Слушаем сообщения из NATS JetStream
func ListenAndProcessMessages(conn db.DBConnection) error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	sub, err := js.Subscribe(natsSubject, func(msg *nats.Msg) {
		var student models.Student
		err := json.Unmarshal(msg.Data, &student)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			return
		}

		// Создаем контекст для каждой горутины
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // Убедитесь, что контекст будет отменен после завершения работы горутины

		// Используем отдельную горутину для сохранения в базе данных
		go func() {
			if err := conn.SaveStudent(ctx, student); err != nil {
				log.Printf("Error saving student to DB: %v", err)
				return
			}
			log.Printf("Successfully saved student: %v", student.FullName)
		}()
	}, nats.ManualAck()) // Устанавливаем ручное подтверждение

	if err != nil {
		return err
	}

	// Обрабатываем сообщения в бесконечном цикле
	for {
		// Ожидаем сообщений в определенный промежуток времени
		msg, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}

		// Подтверждаем обработку сообщения
		msg.Ack()
	}
}
