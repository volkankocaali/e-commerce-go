package database

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestMongoConnection(t *testing.T) {
	mongoUri := "mongodb://127.0.0.1:27019"
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		t.Fatalf("Failed to create mongo connection: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to ping mongo connection: %v", err)
	}

	assert.NoError(t, err, "Failed to create mongo connection")

}
