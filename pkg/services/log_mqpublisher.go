package services

import (
	"github.com/streadway/amqp"
	"log"
)

type LogMQPublisher struct {
	Connection *amqp.Connection
}

func NewLogMQPublisher(conn *amqp.Connection) *LogMQPublisher {
	return &LogMQPublisher{
		Connection: conn,
	}
}

func (lmp *LogMQPublisher) PublishToQueue(message []byte) {
	// Create a channel
	ch, err := lmp.Connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// rabbit mq create queue
	q, err := ch.QueueDeclare(
		"audit_logs_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// rabbit mq publish message
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	log.Println("Published message to queue")
}
