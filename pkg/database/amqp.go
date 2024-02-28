package database

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"log"
)

func NewAmqpClient(cfg config.Config) (*amqp.Connection, error) {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort)
	conn, err := amqp.Dial(amqpUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Kanal açılamadı")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"logQueue", // Kuyruk adı
		true,       // Durable
		false,      // Delete when unused
		false,      // Exclusive
		false,      // No-wait
		nil,        // Arguments
	)
	failOnError(err, "Kuyruk oluşturulamadı")

	logMessage := "Bu bir test log mesajıdır"
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(logMessage),
		})
	failOnError(err, "Log RabbitMQ'ya gönderilemedi")
	log.Printf("Gönderilen log mesajı: %s", logMessage)
	return conn, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
