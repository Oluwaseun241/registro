package kafka

import (
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *ckafka.Producer

func InitKafka(broker string) error {
	var err error
	producer, err = ckafka.NewProducer(&ckafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return err
	}
	log.Println("Kafka producer initialized successfully.")
	return nil
}

func ProduceEvent(topic string, message []byte) error {
	deliveryChan := make(chan ckafka.Event)

	err := producer.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	// Wait for delivery report
	e := <-deliveryChan
	m := e.(*ckafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v", m.TopicPartition.Error)
	} else {
		log.Printf("Delivered message to %v", m.TopicPartition)
	}
	close(deliveryChan)
	return nil
}
