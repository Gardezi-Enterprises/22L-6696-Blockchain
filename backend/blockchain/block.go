package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index      int      `json:"index"`
	Timestamp  string   `json:"timestamp"`
	Data       []string `json:"data"` // Transactions stored as strings
	PrevHash   string   `json:"prevHash"`
	Hash       string   `json:"hash"`
	Nonce      int      `json:"nonce"`
	MerkleRoot string   `json:"merkleRoot"`
	Difficulty int      `json:"difficulty"`
}

// CalculateHash computes the SHA-256 hash for the block
func (b *Block) CalculateHash() string {
	dataBytes, _ := json.Marshal(b.Data)
	record := fmt.Sprintf(
		"%d%s%s%s%s%d",
		b.Index,
		b.Timestamp,
		string(dataBytes),
		b.PrevHash,
		b.MerkleRoot,
		b.Nonce,
	)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// NewBlock creates a new block with the given data and previous hash
func NewBlock(index int, data []string, prevHash string, difficulty int) *Block {
	merkleRoot := BuildMerkleTree(data).Root

	block := &Block{
		Index:      index,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Data:       data,
		PrevHash:   prevHash,
		Nonce:      0,
		MerkleRoot: merkleRoot,
		Difficulty: difficulty,
	}
	return block
}

// CreateGenesisBlock creates the first block in the blockchain (Genesis Block)
func CreateGenesisBlock(difficulty int) *Block {
	genesisData := []string{"Genesis Block - The first block in the chain"}
	merkleRoot := BuildMerkleTree(genesisData).Root

	genesis := &Block{
		Index:      0,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Data:       genesisData,
		PrevHash:   "0000000000000000000000000000000000000000000000000000000000000000",
		Nonce:      0,
		MerkleRoot: merkleRoot,
		Difficulty: difficulty,
	}

	// Mine the genesis block using Proof of Work
	pow := NewProofOfWork(genesis)
	nonce, hash := pow.Run()
	genesis.Nonce = nonce
	genesis.Hash = hash

	return genesis
}
