package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	c, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
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
