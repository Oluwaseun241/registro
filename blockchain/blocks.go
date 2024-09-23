package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Event     string
	PrevHash  string
	Hash      string
}

func (b *Block) CalculateHash() string {
	record := string(b.Index) + b.Timestamp.String() + b.Event + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
