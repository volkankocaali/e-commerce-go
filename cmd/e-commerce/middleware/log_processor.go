package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/volkankocaali/e-commorce-go/pkg/services"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/services/interface"
	"log"
	"time"
)

func LogProcessorMiddleware(ctx *fiber.Ctx, amqpConn *amqp.Connection) error {
	request, err := MinifyJSON(ctx.Request().Body())
	requestHeaders := string(ctx.Request().Header.Header())

	if err != nil {
		log.Println("Failed to minify request:", err)
	}

	// wait for the response
	err = ctx.Next()
	if err != nil {
		return err
	}

	// ctx.Next() is called, so we can access the response now
	response, err := MinifyJSON(ctx.Response().Body())
	responseHeaders := string(ctx.Response().Header.Header())

	if err != nil {
		log.Println("Failed to minify response:", err)
	}

	logMessage := interfaces.LogMessage{
		Request:         request,
		RequestHeaders:  requestHeaders,
		Response:        response,
		ResponseHeaders: responseHeaders,
		CreatedAt:       time.Now().UTC(),
	}

	logBody, err := json.Marshal(logMessage)

	if err != nil {
		log.Println("Failed to marshal log message:", err)
		return err
	}

	publisher := services.NewLogMQPublisher(amqpConn)
	publisher.PublishToQueue(logBody)

	return nil
}

func MinifyJSON(jsonBytes []byte) (string, error) {
	var temp interface{}
	err := json.Unmarshal(jsonBytes, &temp)
	if err != nil {
		return "", err
	}

	minifiedBytes, err := json.Marshal(temp)
	if err != nil {
		return "", err
	}

	return string(minifiedBytes), nil
}
