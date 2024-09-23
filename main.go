package main

import (
	"log"
	"os"
	"registro/handler"
	"registro/kafka"

	"github.com/labstack/echo/v4"
)

func main() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	// Initialize Kafka producer and consumer
	err := kafka.InitKafka(kafkaBroker)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka: %v", err)
	}
	// Start Kafka consumer in a goroutine
	go kafka.StartConsumer()

	e := echo.New()

	e.POST("/produce", handler.ProduceEvent)
	e.GET("/events", handler.GetEvents)
	e.GET("/blockchain", handler.ValidateBlockchain)

	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
