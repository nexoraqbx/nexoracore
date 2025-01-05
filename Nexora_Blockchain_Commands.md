# **Nexora Blockchain Command List**

---

## **1. General Commands**

### 1.1 Query Node and Blockchain Status
- **Query Node Status**:
  ```bash
  curl http://localhost:26657/status
  ```

### 1.2 Query Account and Balances
- **View Account Details**:
  ```bash
  ./build/nexorad query account [account_address]
  ```
- **Query Bank Balances**:
  ```bash
  ./build/nexorad query bank balances [account_address]
  ```

### 1.3 Query Staking and Distribution
- **View Validator Details**:
  ```bash
  ./build/nexorad query staking validators
  ```
- **Query Community Pool**:
  ```bash
  ./build/nexorad query distribution community-pool
  ```
- **Query Validator Outstanding Rewards**:
  ```bash
  ./build/nexorad query distribution validator-outstanding-rewards [validator_address]
  ```

---

## **2. Transaction Commands**

### 2.1 Token Transfers
- **Send Tokens**:
  ```bash
  ./build/nexorad tx bank send [from_address] [to_address] [amount]uqubit \
      --chain-id=nexora-chain \
      --from=[key_name] \
      --fees=1uqubit \
      --yes
  ```

### 2.2 Withdraw Rewards and Commissions
- **Withdraw Validator Commission**:
  ```bash
  ./build/nexorad tx distribution withdraw-rewards [validator_address] \
      --commission \
      --from [key_name] \
      --fees=1uqubit \
      --account-number [account_number] \
      --sequence [sequence] \
      --chain-id=nexora-chain \
      --yes
  ```
- **Withdraw Delegator Rewards**:
  ```bash
  ./build/nexorad tx distribution withdraw-rewards [validator_address] \
      --from [key_name] \
      --fees=1uqubit \
      --account-number [account_number] \
      --sequence [sequence] \
      --chain-id=nexora-chain \
      --yes
  ```

---

## **3. Delegation Commands**

### 3.1 Delegating Tokens
- **Delegate Tokens to Validator**:
  ```bash
  ./build/nexorad tx staking delegate [validator_address] [amount]uqubit \
      --from [key_name] \
      --chain-id=nexora-chain \
      --fees=1uqubit \
      --yes
  ```

### 3.2 Query Delegations
- **Query Delegations to a Validator**:
  ```bash
  ./build/nexorad query staking delegations-to [validator_address]
  ```

---

## **4. Governance Commands**

### 4.1 Query Governance Details
- **List All Proposals**:
  ```bash
  ./build/nexorad query gov proposals
  ```
- **Query Proposal Details**:
  ```bash
  ./build/nexorad query gov proposal [proposal_id]
  ```
- **Query Deposits**:
  ```bash
  ./build/nexorad query gov deposits [proposal_id]
  ```
- **Query Votes**:
  ```bash
  ./build/nexorad query gov votes [proposal_id]
  ```
- **Query Governance Parameters**:
  ```bash
  ./build/nexorad query gov params
  ```
- **Query Tally Results**:
  ```bash
  ./build/nexorad query gov tally [proposal_id]
  ```

### 4.2 Submit and Manage Proposals
- **Submit a Proposal**:
  ```bash
  ./build/nexorad tx gov submit-proposal \
      --title="[Title]" \
      --description="[Description]" \
      --type="Text" \
      --deposit="[amount]uqubit" \
      --from=[proposer_address] \
      --chain-id=nexora-chain \
      --fees=1uqubit \
      --yes
  ```
- **Deposit on a Proposal**:
  ```bash
  ./build/nexorad tx gov deposit [proposal_id] [amount]uqubit \
      --from=[depositor_address] \
      --chain-id=nexora-chain \
      --fees=1uqubit \
      --yes
  ```
- **Vote on a Proposal**:
  ```bash
  ./build/nexorad tx gov vote [proposal_id] [yes|no|abstain|no_with_veto] \
      --from=[voter_address] \
      --chain-id=nexora-chain \
      --fees=1uqubit \
      --yes
  ```

---

## **5. Blockchain Management**

### 5.1 Initialize and Manage the Blockchain
- **Initialize the Blockchain**:
  ```bash
  ./build/nexorad init [moniker_name] --chain-id=nexora-chain
  ```
- **Start the Blockchain Node**:
  ```bash
  ./build/nexorad start
  ```
- **Export Blockchain State**:
  ```bash
  ./build/nexorad export
  ```
- **Rollback Blockchain State**:
  ```bash
  ./build/nexorad rollback
  ```
- **Prune Blockchain State**:
  ```bash
  ./build/nexorad prune
  ```

---

## **6. Logs and Indexing**

### 6.1 Transaction Logs
- **Check Transaction Indexer Configuration**:
  ```bash
  tail -n 20 ~/.nexora/config/config.toml
  ```
- **Enable Detailed Transaction Indexing**:
  ```bash
  # Edit the file ~/.nexora/config/config.toml
  [tx_index]
  index_all_keys = true
  ```

### 6.2 Query Transactions
- **By Hash**:
  ```bash
  curl "http://localhost:26657/tx?hash=0x[tx_hash]"
  ```
- **By Sender**:
  ```bash
  curl "http://localhost:26657/tx_search?query=\"message.sender='[sender_address]'\""
  ```
- **By Height**:
  ```bash
  curl "http://localhost:26657/tx_search?query=\"tx.height>0\""
  ```

---

## **7. Keys Management**

### 7.1 Manage Keys
- **List Keys**:
  ```bash
  ./build/nexorad keys list
  ```
- **Show Key Details**:
  ```bash
  ./build/nexorad keys show [key_name] --bech [acc|val]
  ```
- **Create a New Key**:
  ```bash
  ./build/nexorad keys add [key_name]
  ```

---

