package blockchain

import (
	"errors"
	"strings"
)

// Blockchain holds the chain of blocks and a pool of pending transactions
type Blockchain struct {
	Chain               []*Block `json:"chain"`
	PendingTransactions []string `json:"pendingTransactions"`
	Difficulty          int      `json:"difficulty"`
}

// NewBlockchain initialises a new blockchain with the Genesis block
func NewBlockchain(difficulty int) *Blockchain {
	if difficulty < 1 {
		difficulty = 3
	}
	genesis := CreateGenesisBlock(difficulty)
	return &Blockchain{
		Chain:               []*Block{genesis},
		PendingTransactions: []string{},
		Difficulty:          difficulty,
	}
}

// AddTransaction appends a transaction string to the pending pool
func (bc *Blockchain) AddTransaction(tx string) {
	bc.PendingTransactions = append(bc.PendingTransactions, tx)
}

// MineBlock mines a new block using all pending transactions
// Returns an error if there are no pending transactions
func (bc *Blockchain) MineBlock() (*Block, error) {
	if len(bc.PendingTransactions) == 0 {
		return nil, errors.New("no pending transactions to mine")
	}

	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(
		len(bc.Chain),
		bc.PendingTransactions,
		prevBlock.Hash,
		bc.Difficulty,
	)

	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.Run()
	newBlock.Nonce = nonce
	newBlock.Hash = hash

	bc.Chain = append(bc.Chain, newBlock)
	bc.PendingTransactions = []string{} // reset pending pool
	return newBlock, nil
}

// IsValid validates the integrity of the entire blockchain
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		current := bc.Chain[i]
		previous := bc.Chain[i-1]

		// Recompute hash and compare
		if current.Hash != current.CalculateHash() {
			return false
		}

		// Check linkage
		if current.PrevHash != previous.Hash {
			return false
		}

		// Verify Proof of Work
		pow := NewProofOfWork(current)
		if !pow.Validate() {
			return false
		}
	}
	return true
}

// SearchByData returns all blocks that contain the given query string in any transaction
func (bc *Blockchain) SearchByData(query string) []*Block {
	var results []*Block
	lower := strings.ToLower(query)
	for _, block := range bc.Chain {
		for _, tx := range block.Data {
			if strings.Contains(strings.ToLower(tx), lower) {
				results = append(results, block)
				break
			}
		}
	}
	return results
}

// GetBlockByIndex returns a block at the given index, or nil
func (bc *Blockchain) GetBlockByIndex(index int) *Block {
	if index < 0 || index >= len(bc.Chain) {
		return nil
	}
	return bc.Chain[index]
}

// GetLatestBlock returns the last block in the chain
func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}
