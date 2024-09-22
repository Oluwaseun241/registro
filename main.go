package main

import (
	"log"
	"net/http"
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

	e.GET("/produce", handler.ProduceEvent)
	e.GET("/events", handler.GetEvents)


	log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
