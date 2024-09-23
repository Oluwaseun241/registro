package blockchain

import "time"

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	genesisBook := Block{
		Index:     0,
		Timestamp: time.Now(),
		Event:     "Genesis Block",
		PrevHash:  "",
	}
	genesisBook.Hash = genesisBook.CalculateHash()

	return &Blockchain{Blocks: []Block{genesisBook}}
}

func (bc *Blockchain) AddBlock(event string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Index:     previousBlock.Index + 1,
		Timestamp: time.Now(),
		Event:     event,
		PrevHash:  previousBlock.Hash,
	}

	newBlock.Hash = newBlock.CalculateHash()

	bc.Blocks = append(bc.Blocks, newBlock)
}
