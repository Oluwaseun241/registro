package blockchain

import (
	"errors"
	"fmt"
)

// Ledger represents the blockchain ledger system.
type Ledger struct {
	chain *Blockchain
}

// NewLedger initializes a new ledger with a genesis block.
func NewLedger() *Ledger {
	chain := NewBlockchain()
	return &Ledger{chain: chain}
}

// AddEventToLedger adds a new event to the ledger as a new block.
func (l *Ledger) AddEventToLedger(event string) error {
	if event == "" {
		return errors.New("event data cannot be empty")
	}
	l.chain.AddBlock(event)
	fmt.Printf("New block added to the ledger: %+v\n", l.chain.Blocks[len(l.chain.Blocks)-1])
	return nil
}

// GetBlockchain returns the current state of the blockchain.
func (l *Ledger) GetBlockchain() []Block {
	return l.chain.Blocks
}

// ValidateLedger checks the integrity of the blockchain.
func (l *Ledger) ValidateLedger() bool {
	for i := 1; i < len(l.chain.Blocks); i++ {
		currentBlock := l.chain.Blocks[i]
		previousBlock := l.chain.Blocks[i-1]

		// Check if the current block's previous hash matches the previous block's hash
		if currentBlock.PrevHash != previousBlock.Hash {
			fmt.Println("Blockchain validation failed: Previous hash does not match.")
			return false
		}

		// Check if the current block's hash is correct
		if currentBlock.Hash != currentBlock.CalculateHash() {
			fmt.Println("Blockchain validation failed: Block hash is incorrect.")
			return false
		}
	}
	fmt.Println("Blockchain validation passed.")
	return true
}
