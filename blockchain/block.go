package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	data         map[string]any
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

func (b *Block) calculateHash() string {
	blockData := fmt.Sprintf("%v%v%v%d", b.previousHash, b.data, b.timestamp.String(), b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}
