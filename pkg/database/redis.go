package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
)

func NewRedisClient(cfg config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis")
		return nil, err
	}

	return client, nil
}
