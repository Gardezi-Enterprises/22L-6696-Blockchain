package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	blockchain *Blockchain
	mu         sync.Mutex
	pending    []string // pending transactions waiting to be mined
	pendingMu  sync.Mutex
)

// enableCORS sets CORS headers for the React frontend
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// respondJSON writes a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	enableCORS(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// handleGetBlockchain returns the full blockchain
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	respondJSON(w, http.StatusOK, blockchain)
}

// handleAddTransaction adds a transaction to the pending pool
func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		respondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	var body struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Data == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid transaction data"})
		return
	}

	pendingMu.Lock()
	pending = append(pending, body.Data)
	pendingMu.Unlock()

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"message":     "Transaction added to pending pool",
		"transaction": body.Data,
		"pending":     pending,
	})
}

// handleGetPending returns all pending transactions
func handleGetPending(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}
	pendingMu.Lock()
	defer pendingMu.Unlock()
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"pending": pending,
	})
}

// handleMineBlock mines a new block with pending transactions
func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		respondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	pendingMu.Lock()
	if len(pending) == 0 {
		pendingMu.Unlock()
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "No pending transactions to mine"})
		return
	}

	txs := make([]string, len(pending))
	copy(txs, pending)
	pending = nil
	pendingMu.Unlock()

	mu.Lock()
	newBlock := blockchain.AddBlock(txs)
	mu.Unlock()

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Block mined successfully!",
		"block":   newBlock,
	})
}

// handleSearch searches for data across the blockchain
func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Search query is required"})
		return
	}

	mu.Lock()
	results := blockchain.SearchBlockchain(query)
	mu.Unlock()

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"query":   query,
		"results": results,
		"count":   len(results),
	})
}

func main() {
	// Create the blockchain named after the student
	blockchain = NewBlockchain("SyedTaseerHaider")
	fmt.Println("SyedTaseerHaider Blockchain initialized!")
	fmt.Printf("Genesis Block Hash: %s\n", blockchain.Blocks[0].Hash)

	// API Routes
	http.HandleFunc("/api/blockchain", handleGetBlockchain)
	http.HandleFunc("/api/transaction", handleAddTransaction)
	http.HandleFunc("/api/pending", handleGetPending)
	http.HandleFunc("/api/mine", handleMineBlock)
	http.HandleFunc("/api/search", handleSearch)

	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
