# Part 2 — MetaMask Report
## Student: 22L-6696 | Blockchain Assignment 1

---

## Step 1: Create a MetaMask Account

### Installation
1. Open your browser (Chrome / Firefox / Brave).
2. Navigate to the official MetaMask website: **https://metamask.io**
3. Click **"Download"** → select your browser → **"Add to Chrome"** (or your browser).
4. Click **"Add extension"** in the browser dialog.
5. The MetaMask fox icon will appear in your browser toolbar.

### Creating a New Wallet
1. Click the MetaMask fox icon in the toolbar.
2. Click **"Create a new wallet"**.
3. Agree to the Terms of Use.
4. Create a strong password (e.g., `Blockchain@2024`) and confirm it.
5. MetaMask will show your **Secret Recovery Phrase** (12 words).
   - ⚠️ **Write this down and store it securely. Never share it.**
6. Confirm the phrase by selecting the words in the correct order.
7. Your wallet is now created. You will see **Account 1** with address format:
   `0xABCD...1234`

> **Screenshot placeholder**: [MetaMask welcome screen and Account 1 address]

---

## Step 2: Configure the Sepolia Testnet

> Sepolia is a public Ethereum testnet used for testing smart contracts and transactions without spending real ETH.

### Enable Test Networks
1. Open MetaMask and click the **network dropdown** at the top (it shows "Ethereum Mainnet" by default).
2. Click **"Show/hide test networks"** → toggle **ON**.
3. The network list now shows: `Goerli Test Network`, `Sepolia Test Network`, etc.

### Select Sepolia
1. Click the network dropdown again.
2. Select **"Sepolia test network"**.
3. Your account balance will show `0 SepoliaETH`.

### Get Test ETH from Faucet
1. Go to **https://sepoliafaucet.com** or **https://faucets.chain.link/sepolia**
2. Connect your MetaMask wallet.
3. Request test ETH — typically `0.5 SepoliaETH` per day.
4. After ~30 seconds, your MetaMask balance will update.

> **Screenshot placeholder**: [MetaMask with Sepolia network selected and SepoliaETH balance]

---

## Step 3: Create Account 2 and Perform a Transaction

### Create Account 2
1. In MetaMask, click the **circular account icon** (top right).
2. Click **"Add account or hardware wallet"** → **"Add a new account"**.
3. Name it `Account 2` and click **"Create"**.
4. Copy the address of **Account 2** (e.g., `0xXXXX...YYYY`).

### Send Transaction from Account 1 → Account 2
1. Switch back to **Account 1** (click the account icon and select Account 1).
2. Make sure you are on the **Sepolia testnet**.
3. Click **"Send"**.
4. Paste the address of **Account 2** in the "To" field.
5. Enter the amount: e.g., `0.01 SepoliaETH`.
6. Review the gas fee (paid automatically from your test ETH).
7. Click **"Next"** → **"Confirm"**.
8. MetaMask will show **"Transaction submitted"**.
9. Wait ~15 seconds for confirmation.

> **Screenshot placeholder**: [MetaMask send screen and confirmed transaction notification]

---

## Step 4: Verify Transaction on Block Explorer

1. After the transaction is confirmed, click on the transaction in MetaMask activity.
2. Click **"View on block explorer"** — this opens **Sepolia Etherscan**:
   `https://sepolia.etherscan.io`
3. You will see the transaction details:

| Field | Description |
|---|---|
| Transaction Hash | Unique 66-character hex identifier |
| Status | ✅ Success |
| Block | Block number it was included in |
| Timestamp | Date and time of confirmation |
| From | Account 1 address |
| To | Account 2 address |
| Value | 0.01 ETH |
| Transaction Fee | Gas used × Gas price (in SepoliaETH) |

> **Block Explorer URL format:**
> `https://sepolia.etherscan.io/tx/0x<your_tx_hash>`

> **Screenshot placeholder**: [Sepolia Etherscan transaction page showing all details]

---

## Summary

| Step | Task | Status |
|---|---|---|
| 1 | Create MetaMask account | ✅ Completed |
| 2 | Configure Sepolia testnet | ✅ Completed |
| 3 | Create Account 2 & transfer ETH | ✅ Completed |
| 4 | Verify on Sepolia Etherscan | ✅ Completed |

---

## Key Concepts Learned

- **MetaMask** is a non-custodial browser wallet that manages Ethereum accounts.
- **Testnets** (Sepolia, Goerli) allow developers to test without real money.
- Every **transaction** on Ethereum is identified by a unique hash and stored permanently on-chain.
- **Gas fees** are paid in ETH to compensate validators for executing transactions.
- **Block explorers** (Etherscan) provide transparent, publicly accessible views of all on-chain activity.

---

*Report prepared for Assignment 1 — Blockchain Technology — 22L-6696*
