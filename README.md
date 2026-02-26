# Blockchain Project — 22L-6696 A1

A full-stack Blockchain implementation in **Go** (backend) and **React** (frontend).

---

## Project Structure

```
22L-6696 A1 blockchain/
├── backend/                  # Go blockchain API server
│   ├── main.go               # HTTP server & REST endpoints
│   ├── go.mod
│   └── blockchain/
│       ├── block.go          # Block struct & Genesis block
│       ├── blockchain.go     # Chain management
│       ├── merkle.go         # Merkle Tree implementation
│       └── pow.go            # Proof of Work algorithm
├── frontend/                 # React UI
│   ├── public/index.html
│   ├── package.json
│   └── src/
│       ├── index.js
│       ├── App.js
│       ├── App.css
│       └── components/
│           ├── ViewBlockchain.js
│           ├── AddTransaction.js
│           ├── MineBlock.js
│           ├── PendingTransactions.js
│           └── SearchBlock.js
└── README.md
```

---

## Features

| Feature | Description |
|---|---|
| Block Structure | Index, Timestamp, Data (Transactions), PrevHash, Hash, Nonce, MerkleRoot, Difficulty |
| Genesis Block | Automatically created on server startup |
| Merkle Tree | Built from transactions; root stored in each block |
| Transactions | Add string transactions to a pending pool |
| Proof of Work | SHA-256 mining with configurable difficulty (default: 3 leading zeros) |
| View Blockchain | See all blocks with expandable details |
| Search | Search transaction data across the entire chain |
| React UI | Full interactive frontend |

---

## Running the Project

### Prerequisites
- [Go 1.21+](https://golang.org/dl/)
- [Node.js 18+](https://nodejs.org/)

### 1. Start the Go Backend

```bash
cd backend
go run main.go
```

The server starts on **http://localhost:8080**

### 2. Start the React Frontend

```bash
cd frontend
npm install
npm start
```

The app opens on **http://localhost:3000**

---

## REST API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/chain` | Return the full blockchain |
| GET | `/pending` | Return pending transactions |
| POST | `/transaction` | Add a transaction `{ "transaction": "..." }` |
| POST | `/mine` | Mine a new block from pending transactions |
| GET | `/search?q=query` | Search blocks by transaction content |
| GET | `/validate` | Validate blockchain integrity |

---

## How It Works

### Block Structure
```go
type Block struct {
    Index      int      // Block height
    Timestamp  string   // ISO timestamp
    Data       []string // Transactions
    PrevHash   string   // Hash of previous block
    Hash       string   // SHA-256 hash of this block
    Nonce      int      // Mining nonce
    MerkleRoot string   // Merkle root of transactions
    Difficulty int      // PoW difficulty
}
```

### Proof of Work
The miner increments the nonce until:
```
SHA256(index + timestamp + data + prevHash + merkleRoot + nonce)
```
produces a hash starting with `difficulty` number of leading zeros.

### Merkle Tree
Transactions are hashed into leaf nodes. Pairs are recursively combined (SHA-256 of concatenated hashes) up to the root. The root is stored in the block and can be used to verify individual transactions without downloading the full block.
