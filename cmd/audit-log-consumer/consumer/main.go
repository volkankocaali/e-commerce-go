package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/streadway/amqp"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	"github.com/volkankocaali/e-commorce-go/pkg/services"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/services/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize MongoDB and AMQP clients
	mongoDBClient, err := database.NewMongoClient(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize AMQP client
	amqpClient, err := database.NewAMQPClient(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to AMQP: %v", err)
	}

	// Setup AMQP channel and queue
	amqpChannel, err := setupAMQPChannel(amqpClient)
	if err != nil {
		log.Fatalf("Failed to setup AMQP channel and queue: %v", err)
	}
	defer amqpChannel.Close()

	// Start listening for messages
	listenForMessages(amqpChannel, mongoDBClient)
}

func setupAMQPChannel(amqpClient *amqp.Connection) (*amqp.Channel, error) {
	channel, err := amqpClient.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	_, err = channel.QueueDeclare(
		"audit_logs_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return channel, nil
}

func listenForMessages(channel *amqp.Channel, mongoDBClient *mongo.Client) {
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")

	msgs, err := channel.Consume(
		"audit_logs_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	mongoLogger, err := services.NewMongoDBLogger(mongoDBClient)
	if err != nil {
		log.Fatalf("Failed to create MongoDB logger: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			go processMessage(msg, mongoLogger)
		}
	}()

	<-forever
}

func processMessage(msg amqp.Delivery, logger *services.MongoDBLogger) {
	var logMessage interfaces.LogMessage
	if err := json.Unmarshal(msg.Body, &logMessage); err != nil {
		log.Fatalf("Failed to unmarshal message: %v", err)
		return
	}

	if err := logger.Process(logMessage); err != nil {
		log.Fatalf("Failed to log to MongoDB: %v", err)
	}
}
