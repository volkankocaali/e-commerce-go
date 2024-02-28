package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppPort string
	AppName string

	MongoDBURI string

	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	RedisHost     string
	RedisPort     string
	RedisPassword string

	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassword string

	ElasticSearchHost string
	ElasticSearchPort string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),

		MongoDBURI: os.Getenv("MONGODB_SERVER"),

		MySQLHost:     os.Getenv("MYSQL_HOST"),
		MySQLPort:     os.Getenv("MYSQL_PORT"),
		MySQLUser:     os.Getenv("MYSQL_USER"),
		MySQLPassword: os.Getenv("MYSQL_PASSWORD"),
		MySQLDatabase: os.Getenv("MYSQL_DATABASE"),

		RedisHost:     os.Getenv("REDIS_SERVER_HOST"),
		RedisPort:     os.Getenv("REDIS_SERVER_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),

		RabbitMQHost:     os.Getenv("RABBITMQ_SERVER_HOST"),
		RabbitMQPort:     os.Getenv("RABBITMQ_SERVER_PORT"),
		RabbitMQUser:     os.Getenv("RABBITMQ_SERVER_USER"),
		RabbitMQPassword: os.Getenv("RABBITMQ_SERVER_PASSWORD"),

		ElasticSearchHost: os.Getenv("ELASTICSEARCH_SERVER_HOST"),
		ElasticSearchPort: os.Getenv("ELASTICSEARCH_SERVER_PORT"),
	}

	return cfg
}
