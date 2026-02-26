package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleNode represents a node in the Merkle Tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  string
}

// MerkleTree represents the full Merkle Tree
type MerkleTree struct {
	Root *MerkleNode
}

// hashData hashes a string using SHA-256
func hashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// NewMerkleNode creates a new Merkle Tree node
func NewMerkleNode(left, right *MerkleNode, data string) *MerkleNode {
	node := &MerkleNode{}

	if left == nil && right == nil {
		// Leaf node
		node.Hash = hashData(data)
	} else {
		// Internal node: hash of children's hashes
		combinedHash := ""
		if left != nil {
			combinedHash += left.Hash
		}
		if right != nil {
			combinedHash += right.Hash
		}
		node.Hash = hashData(combinedHash)
	}

	node.Left = left
	node.Right = right
	return node
}

// NewMerkleTree builds a Merkle Tree from a slice of transaction strings
func NewMerkleTree(transactions []string) *MerkleTree {
	if len(transactions) == 0 {
		// If no transactions, create a tree with an empty hash
		root := NewMerkleNode(nil, nil, "")
		return &MerkleTree{Root: root}
	}

	var nodes []*MerkleNode

	// Create leaf nodes from transactions
	for _, tx := range transactions {
		node := NewMerkleNode(nil, nil, tx)
		nodes = append(nodes, node)
	}

	// Build tree bottom-up
	for len(nodes) > 1 {
		var level []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				node := NewMerkleNode(nodes[i], nodes[i+1], "")
				level = append(level, node)
			} else {
				// Odd number of nodes: duplicate the last one
				node := NewMerkleNode(nodes[i], nodes[i], "")
				level = append(level, node)
			}
		}

		nodes = level
	}

	return &MerkleTree{Root: nodes[0]}
}

// GetMerkleRoot returns the root hash of the Merkle Tree
func GetMerkleRoot(transactions []string) string {
	tree := NewMerkleTree(transactions)
	return tree.Root.Hash
}
