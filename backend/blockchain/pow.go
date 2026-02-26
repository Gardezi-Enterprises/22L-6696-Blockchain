package blockchain

import (
	"fmt"
	"strings"
)

// ProofOfWork holds a reference to the block being mined
type ProofOfWork struct {
	Block *Block
}

// NewProofOfWork creates a new ProofOfWork instance for a given block
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{Block: block}
}

// target returns the difficulty target prefix (e.g., "0000" for difficulty 4)
func (pow *ProofOfWork) target() string {
	return strings.Repeat("0", pow.Block.Difficulty)
}

// Validate checks whether the block's hash satisfies the difficulty requirement
func (pow *ProofOfWork) Validate() bool {
	return strings.HasPrefix(pow.Block.Hash, pow.target())
}

// Run performs the proof-of-work mining loop
// It iterates through nonce values until a valid hash is found
func (pow *ProofOfWork) Run() (int, string) {
	nonce := 0
	target := pow.target()

	fmt.Printf("Mining block %d with difficulty %d ...\n", pow.Block.Index, pow.Block.Difficulty)

	for {
		pow.Block.Nonce = nonce
		hash := pow.Block.CalculateHash()

		if strings.HasPrefix(hash, target) {
			fmt.Printf("Block %d mined! Nonce: %d  Hash: %s\n", pow.Block.Index, nonce, hash)
			return nonce, hash
		}
		nonce++
	}
}
