package kafka

import (
	"context"
	"log"

	ckafka "github.com/segmentio/kafka-go"
)

var producer *ckafka.Writer

func InitKafka(broker string) error {
	producer = ckafka.NewWriter(ckafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   "libro-events",
		Async:   true,
	})
	log.Println("Kafka producer initialized successfully.")
	return nil
}

func ProduceEvent(topic string, message []byte) error {
	producer.Topic = topic

	err := producer.WriteMessages(context.Background(), ckafka.Message{
		Value: message,
	})

	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	log.Println("Message delivered successfully.")
	return nil
}
