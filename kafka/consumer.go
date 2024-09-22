package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func StartConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "libro-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{"libro-events"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received event: %s\n", string(msg.Value))
			// TODO: Add logic to record to blockchain
		} else {
			log.Printf("Error reading message: %v", err)
		}
	}
}
