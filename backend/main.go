package main

import (
	"blockchain-api/blockchain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Global blockchain instance (difficulty = 3 leading zeros)
var bc = blockchain.NewBlockchain(3)

// ---- CORS middleware ----
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// ---- Response helpers ----
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func errorResponse(w http.ResponseWriter, status int, msg string) {
	jsonResponse(w, status, map[string]string{"error": msg})
}

// ---- Handlers ----

// GET /chain – return the full blockchain
func handleGetChain(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"chain":  bc.Chain,
		"length": len(bc.Chain),
		"valid":  bc.IsValid(),
	})
}

// GET /pending – return pending transactions
func handleGetPending(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"pendingTransactions": bc.PendingTransactions,
		"count":               len(bc.PendingTransactions),
	})
}

// POST /transaction – add a new transaction
// Body: { "transaction": "Alice sends Bob 10 BTC" }
func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Transaction string `json:"transaction"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || strings.TrimSpace(body.Transaction) == "" {
		errorResponse(w, http.StatusBadRequest, "invalid or missing 'transaction' field")
		return
	}

	bc.AddTransaction(body.Transaction)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"message":             "Transaction added to pending pool",
		"transaction":         body.Transaction,
		"pendingTransactions": bc.PendingTransactions,
	})
}

// POST /mine – mine a new block from pending transactions
func handleMine(w http.ResponseWriter, r *http.Request) {
	block, err := bc.MineBlock()
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Block %d mined successfully!", block.Index),
		"block":   block,
	})
}

// GET /search?q=query – search for blocks containing the query string
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		errorResponse(w, http.StatusBadRequest, "query parameter 'q' is required")
		return
	}

	results := bc.SearchByData(query)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"query":   query,
		"results": results,
		"count":   len(results),
	})
}

// GET /validate – check if the blockchain is valid
func handleValidate(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"valid":   bc.IsValid(),
		"message": "Blockchain integrity check complete",
	})
}

// ---- Main ----
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/chain", enableCORS(handleGetChain))
	mux.HandleFunc("/pending", enableCORS(handleGetPending))
	mux.HandleFunc("/transaction", enableCORS(handleAddTransaction))
	mux.HandleFunc("/mine", enableCORS(handleMine))
	mux.HandleFunc("/search", enableCORS(handleSearch))
	mux.HandleFunc("/validate", enableCORS(handleValidate))

	port := ":8080"
	fmt.Printf("Blockchain API server running on http://localhost%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /chain       - View the full blockchain")
	fmt.Println("  GET  /pending     - View pending transactions")
	fmt.Println("  POST /transaction - Add a transaction")
	fmt.Println("  POST /mine        - Mine a new block")
	fmt.Println("  GET  /search?q=xx - Search the blockchain")
	fmt.Println("  GET  /validate    - Validate the chain")

	log.Fatal(http.ListenAndServe(port, mux))
}
