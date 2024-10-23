package nats

import (
	"ape/internal/db"
	"ape/internal/models"
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

	sub, err := js.SubscribeSync(natsSubject)
	if err != nil {
		return err
	}

	for {
		msg, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}

		var student models.Student
		err = json.Unmarshal(msg.Data, &student)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			continue
		}

		err = conn.SaveStudent(student)
		if err != nil {
			log.Printf("Error saving student to DB: %v", err)
			continue
		}

		log.Printf("Successfully saved student: %v", student.FullName)
	}
}
