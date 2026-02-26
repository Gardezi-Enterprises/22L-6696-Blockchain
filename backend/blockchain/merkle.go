package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleNode represents a node in the Merkle tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  string
}

// MerkleTree represents the full Merkle tree
type MerkleTree struct {
	Root     string
	RootNode *MerkleNode
}

// hashData returns the SHA-256 hash of the given data
func hashData(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// newMerkleNode creates a new Merkle tree node
func newMerkleNode(left, right *MerkleNode, data string) *MerkleNode {
	node := &MerkleNode{}

	if left == nil && right == nil {
		// Leaf node: hash the raw data
		node.Data = hashData(data)
	} else {
		// Internal node: hash the concatenation of children hashes
		combined := ""
		if left != nil {
			combined += left.Data
		}
		if right != nil {
			combined += right.Data
		}
		node.Data = hashData(combined)
	}

	node.Left = left
	node.Right = right
	return node
}

// BuildMerkleTree builds a Merkle tree from a list of transactions and returns it
func BuildMerkleTree(transactions []string) *MerkleTree {
	if len(transactions) == 0 {
		// Empty tree
		node := newMerkleNode(nil, nil, "empty")
		return &MerkleTree{Root: node.Data, RootNode: node}
	}

	// Build leaf nodes
	var nodes []*MerkleNode
	for _, tx := range transactions {
		node := newMerkleNode(nil, nil, tx)
		nodes = append(nodes, node)
	}

	// If odd number of nodes, duplicate the last one
	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}

	// Build tree bottom-up
	for len(nodes) > 1 {
		var level []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			var right *MerkleNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			} else {
				right = left // duplicate if odd
			}
			parent := newMerkleNode(left, right, "")
			level = append(level, parent)
		}
		nodes = level

		// If odd, duplicate the last node at this level
		if len(nodes) > 1 && len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
	}

	return &MerkleTree{Root: nodes[0].Data, RootNode: nodes[0]}
}

// GetMerkleProof returns the proof path for a given transaction index
func GetMerkleProof(transactions []string, targetIndex int) []string {
	if len(transactions) == 0 || targetIndex >= len(transactions) {
		return []string{}
	}

	var nodes []string
	for _, tx := range transactions {
		nodes = append(nodes, hashData(tx))
	}

	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}

	var proof []string
	idx := targetIndex

	for len(nodes) > 1 {
		var nextLevel []string
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i]
			if i+1 < len(nodes) {
				right = nodes[i+1]
			}

			// If current index is part of this pair, add sibling to proof
			if i == idx || i+1 == idx {
				if i == idx {
					proof = append(proof, right) // sibling is right
				} else {
					proof = append(proof, left) // sibling is left
				}
				idx = i / 2
			}

			nextLevel = append(nextLevel, hashData(left+right))
		}
		nodes = nextLevel
		if len(nodes) > 1 && len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
	}

	return proof
}
