package database

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	redisOptions := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(redisOptions)

	_, err := client.Ping().Result()

	if err != nil {
		t.Fatalf("Failed to create redis connection: %v", err)
	}

	assert.NoError(t, err, "Failed to create redis connection")
}
