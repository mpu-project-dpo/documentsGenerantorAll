package nats

import (
	"ape/internal/models"
	"encoding/json"
	"log"

	"github.com/spf13/viper"

	"github.com/nats-io/nats.go"
)

const natsSubject = "students_data"

// ListenAndProcessMessages слушает сообщения из NATS и отправляет их в канал
func ListenAndProcessMessages(c chan<- models.Student) error {
	// Читаем URL из конфигурации
	natsURL := viper.GetString("nats.url")

	// Подключаемся к NATS
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	// Создаем канал для получения сообщений
	msgChan := make(chan *nats.Msg, 64)

	// Подписываемся на канал
	sub, err := nc.ChanSubscribe(natsSubject, msgChan)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	// Обрабатываем сообщения из канала
	for msg := range msgChan {
		var student models.Student
		if err := json.Unmarshal(msg.Data, &student); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Отправляем данные в канал
		c <- student
	}

	return nil
}
