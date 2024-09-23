package blockchain

import "testing"

func TestAddEventToLedger(t *testing.T) {
	ledger := NewLedger()

	event := "Test Event"
	err := ledger.AddEventToLedger(event)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	chain := ledger.GetChain()

	if len(chain) != 2 {
		t.Fatalf("Expected ledger length of 1, got %d", len(chain))
	}

	if chain[0].Event != event {
		t.Fatalf("Expected event data '%s', got '%s'", event, chain[0].Event)
	}
}

func TestValidateLedger(t *testing.T) {
	ledger := NewLedger()
	ledger.AddEventToLedger("Event 1")
	ledger.AddEventToLedger("Event 2")

	if !ledger.ValidateLedger() {
		t.Fatal("Expected ledger to be valid")
	}

	chain := ledger.GetChain()

	// Tamper with the ledger
	chain[1].Event = "Tampered Event"

	if ledger.ValidateLedger() {
		t.Fatal("Expected ledger to be invalid after tampering")
	}
}
