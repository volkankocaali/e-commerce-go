package database

import (
	"context"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

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

	return client, nil
}
