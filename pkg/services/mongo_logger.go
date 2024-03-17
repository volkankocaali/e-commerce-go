package services

import (
	"context"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/services/interface"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoDBLogger struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBLogger(client *mongo.Client) (*MongoDBLogger, error) {
	db := client.Database(database.DatabaseName)
	collection := db.Collection(database.CollectionName)

	return &MongoDBLogger{
		client:     client,
		collection: collection,
	}, nil
}

func (l *MongoDBLogger) Process(message interfaces.LogMessage) error {
	_, err := l.collection.InsertOne(context.Background(), message)

	if err != nil {
		log.Println("Failed to insert into MongoDB:", err)
	}

	return err
}
