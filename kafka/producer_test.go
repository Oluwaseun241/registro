package kafka

import "testing"

func TestInitKafka(t *testing.T) {
	err := InitKafka("localhost:9092")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestProduceEvent(t *testing.T) {
	err := InitKafka("localhost: 9092")
	if err != nil {
		t.Fatalf("Failed to initialize Kafka: %v", err)
	}

	err = ProduceEvent("libro-events", []byte("Test message"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
