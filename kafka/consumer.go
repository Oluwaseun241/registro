package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"registro/blockchain"

	"github.com/segmentio/kafka-go"
	ckafka "github.com/segmentio/kafka-go"
)

var Ledger *blockchain.Ledger

func StartConsumer() {

	Ledger = blockchain.NewLedger()

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	reader := ckafka.NewReader(ckafka.ReaderConfig{
		Brokers:     []string{kafkaBroker},
		GroupID:     "libro-group",
		Topic:       "libro-events",
		StartOffset: kafka.FirstOffset,
	})

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		event := string(msg.Value)
		fmt.Printf("Received event: %s\n", string(msg.Value))

		err = Ledger.AddEventToLedger(event)
		if err != nil {
			log.Printf("Failed to add event to ledger: %v", err)
		}
	}
}
