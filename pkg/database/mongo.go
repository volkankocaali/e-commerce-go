package database

import (
	"context"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	DatabaseName    = "e-commerce"
	CollectionName  = "audit_logs"
	TTLIndexName    = "createdAt"
	TTLIndexSeconds = 60 * 60 * 24 * 5
)

type MongoDBIndexManager struct {
	Client *mongo.Client
}

func NewMongoDBIndexManager(client *mongo.Client) *MongoDBIndexManager {
	return &MongoDBIndexManager{
		Client: client,
	}
}

func (m *MongoDBIndexManager) SetupIndexes() error {
	ctx := context.Background()
	collection := m.Client.Database(DatabaseName).Collection(CollectionName)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: TTLIndexName, Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(TTLIndexSeconds), // 5 days,
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		log.Printf("Failed to create TTL index: %s", err)
		return err
	}

	return nil
}

func NewMongoClient(cfg config.Config) (*mongo.Client, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(cfg.MongoDBURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("mongo connect error : %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("mongo ping error : %s", err)
	}

	// mongo index manager setup
	indexManager := NewMongoDBIndexManager(client)
	err = indexManager.SetupIndexes()
	if err != nil {
		log.Fatalf("Failed to setup indexes: %s", err)
	}

	return client, nil
}
