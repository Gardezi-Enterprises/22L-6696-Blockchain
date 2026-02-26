package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Difficulty defines how many leading zeros the hash must have
const Difficulty = 4

// Block represents a single block in the blockchain
type Block struct {
	Index        int      `json:"index"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
	PrevHash     string   `json:"prevHash"`
	Hash         string   `json:"hash"`
	Nonce        int      `json:"nonce"`
	MerkleRoot   string   `json:"merkleRoot"`
}

// Blockchain holds the chain name and the chain of blocks
type Blockchain struct {
	Name   string   `json:"name"`
	Blocks []*Block `json:"blocks"`
}

// CalculateHash computes the SHA-256 hash of a block
func (b *Block) CalculateHash() string {
	data := fmt.Sprintf("%d%s%s%s%d", b.Index, b.Timestamp, b.MerkleRoot, b.PrevHash, b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// MineBlock performs Proof of Work mining on the block
func (b *Block) MineBlock() {
	target := strings.Repeat("0", Difficulty)
	for {
		b.Hash = b.CalculateHash()
		if strings.HasPrefix(b.Hash, target) {
			break
		}
		b.Nonce++
	}
	fmt.Printf("Block %d mined! Nonce: %d, Hash: %s\n", b.Index, b.Nonce, b.Hash)
}

// NewBlockchain creates a new blockchain with name and genesis block
func NewBlockchain(name string) *Blockchain {
	genesisBlock := createGenesisBlock()
	return &Blockchain{
		Name:   name,
		Blocks: []*Block{genesisBlock},
	}
}

// createGenesisBlock creates the first block of the blockchain
func createGenesisBlock() *Block {
	transactions := []string{"l226696"}
	merkleRoot := GetMerkleRoot(transactions)

	block := &Block{
		Index:        0,
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: transactions,
		PrevHash:     "0",
		Nonce:        0,
		MerkleRoot:   merkleRoot,
	}

	block.MineBlock()
	return block
}

// GetLatestBlock returns the last block in the chain
func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

// AddBlock creates new block with pending transactions and mines it
func (bc *Blockchain) AddBlock(transactions []string) *Block {
	prevBlock := bc.GetLatestBlock()
	merkleRoot := GetMerkleRoot(transactions)

	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Nonce:        0,
		MerkleRoot:   merkleRoot,
	}

	newBlock.MineBlock()
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// SearchBlockchain searches for a string across all blocks' transactions
func (bc *Blockchain) SearchBlockchain(query string) []*Block {
	var results []*Block
	queryLower := strings.ToLower(query)

	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if strings.Contains(strings.ToLower(tx), queryLower) {
				results = append(results, block)
				break
			}
		}
	}
	return results
}
