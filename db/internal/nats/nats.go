package nats

import (
	"encoding/json"
	"log"

	"github.com/mpu-project-dpo/documentsGenerantorAll/pkg/nats-contracts"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

const natsSubject = "students_data"

// ListenAndProcessMessages слушает сообщения из NATS и отправляет их в канал
func ListenAndProcessMessages(c chan<- nats_contracts.Document) error {
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
		var document nats_contracts.Document
		if err := json.Unmarshal(msg.Data, &document); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		c <- document
	}

	return nil
}
