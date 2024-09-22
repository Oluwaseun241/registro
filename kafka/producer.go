package kafka

import (

	"log"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var producer *kafka.Producer

func InitKafka(broker string) error {
	_, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return err
	}
	return nil
}

func ProduceEvent(topic string, message []byte) error {
	deliveryChan := make(chan kafka.Event)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: message,
	}, deliveryChan)
	if err != nil {
        log.Printf("Failed to send message: %v", err)
        return err
    }

    // Wait for delivery report
    e := <-deliveryChan
    m := e.(*kafka.Message)
    if m.TopicPartition.Error != nil {
        log.Printf("Delivery failed: %v", m.TopicPartition.Error)
    } else {
        log.Printf("Delivered message to %v", m.TopicPartition)
    }
    close(deliveryChan)
    return nil
}
