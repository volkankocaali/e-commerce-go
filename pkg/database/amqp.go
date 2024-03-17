package database

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"log"
)

func NewAMQPClient(cfg config.Config) (*amqp.Connection, error) {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
	)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}

	return conn, nil
}
