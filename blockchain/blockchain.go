// Package blockchain implements an example blockchain.
// Based on tutorial: https://blog.logrocket.com/build-blockchain-with-go/
package blockchain

import (
	"errors"
	"time"
)

type Blockchain struct {
	genesisBlock *Block
	chain        []*Block
	difficulty   int
}

// New creates a new blockchain with
// an initial genesis block.
func New(difficulty int) *Blockchain {
	genesis := &Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return &Blockchain{
		genesisBlock: genesis,
		chain:        []*Block{genesis},
		difficulty:   difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) error {
	if len(b.chain) < 1 {
		return errors.New("malformed blockchain - missing genesis block")
	}
	blockData := map[string]any{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	newBlock := &Block{
		data:         blockData,
		previousHash: b.chain[len(b.chain)-1].hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
	return nil
}

func (b *Blockchain) isValid() bool {
	if len(b.chain) < 1 {
		return false
	}
	for i := range b.chain[1:] {
		prev := b.chain[i]
		curr := b.chain[i+1]
		if curr.hash != curr.calculateHash() || curr.previousHash != prev.hash {
			return false
		}
	}
	return true
}
