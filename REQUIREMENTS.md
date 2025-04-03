# Wallet Client App

A wallet client application that works with two cryptocurrencies' coins (referred to as "currencies") and uses blockchain as the server side.
Thus blockchain should be multicurrency

---

## Key Pair

To send and receive coins, a user needs at least one key pair for each currency.  
Each key pair proves ownership of one wallet address for a specific currency.

---

## Currencies

### Questions:
<!-- - How can they be integrated into existing blockchain? -->
<!-- - Is it possible to mint them directly on the blockchain?
  - If so, they might not need to be created elsewhere. -->

Chosen one
Option 1: Bring multicurrency to a blockchain
    currency can be created directly on the blockchain
    currency can be native or non-native to a blockchain
    currency serves as a transfer of money

<!-- Complicated
Option 2: currency created outside of blockchain and transferred over the network
    Need bridge or something like that to integrate two blockchains
    My blockchain should be public to interact with other -->

---

## Blockchain (Server Side)

Used as an **irreversible digital ledger** to:
- Make transactions (TXs)  
- Store transaction history  
- Provide transaction history  

---

## Wallet (Client Side)

### Features:
1. **Basic Authentication** (token or password)  
   - Where should credentials be stored?  
2. **Create Wallet**  
   - The wallet has two pockets, one for each currency.  
   - Optionally, users can create additional pockets.  
   - Each pocket corresponds to a key pair.  
   - Each pocket starts with **X value** of currency's coins (e.g., airdrop or distribution).  
3. **Delete Wallet**  
   - Optionally, users can delete created pockets, but at least one per currency must remain.  
4. **Send Coins**  
   - How to differentiate between currencies when sending?  
   - Incorrect transactions should be reverted to prevent loss of funds.  
5. **Receive Coins**  
6. **Retrieve Data**  
   - Balance of each currency (optionally compare to current BTC price).  
   - Transaction history for each currency over a specific period.  

---

## User Story (Test Scenario)

### **User 1**
1. Start the app.  
2. Register.  
3. Create a wallet.  
4. Get balance of each pocket:  
   - Each pocket contains **X value** of currency's coins.  
5. Get transaction history:  
   - It's empty.  

### **User 2**
6. Start the app.  
7. Register.  
8. Create a wallet.  
9. Get balance of each pocket:  
   - Each pocket contains **X value** of currency's coins.  
10. Get transaction history:  
    - It's empty.  
11. Send **1 unit** of `currency's coins_1` to User 1.  
12. Get balance of pocket with sent currency's coins:  
    - `Balance = previous balance - 1`  
13. Get transaction history:  
    - One transaction (sent to User 1).  

### **User 1**
14. Get balance of pocket with received currency's coins:  
    - `Balance = previous balance + 1`  
15. Get transaction history:  
    - One received transaction.  
16. Send **1 unit** of `currency's coins_2` to User 2.  
17. Get balance of pocket with sent currency's coins:  
    - `Balance = previous balance - 1`  
18. Get transaction history:  
    - One sent transaction.  
    - One received transaction.  

### **User 2**
19. Get balance of pocket with received currency's coins:  
    - `Balance = previous balance + 1`  
20. Get transaction history:  
    - One sent transaction.  
    - One received transaction.  

### **User 1 & User 2**
21. Delete their wallets.  
    - **Coins are gone forever.**  

## Architecture
### Bring multicurrency to the blockchain
Converting a single to a multi-currency blockchain/ledger requires additional modules:
- Multi-Coins/Tokens Ledger Design
    Coins: chrome-extension://efaidnbmnnnibpcajpcglclefindmkaj/https://eprint.iacr.org/2020/895.pdf
    Tokens: https://xrpl.org/
    A ledger is assumed to be a list of transactions

- Multi-currency transaction processing
- Token issuance 

- Modular architecture
- Multi-Currency Wallet Integration
- Consensus and Validation
    Adjustments to the consensus mechanism may be necessary if the introduction of multiple currencies impacts transaction throughput or security assumptions. In some designs, you might separate the validation of different asset types or incorporate additional verification steps for token-specific transactions.

#### Multi-token ledger design
#### Multi-currency transaction processing
#### Token issuance 


### Make wallet multicurrency
### Wallet UI

## Projecting
### Bring multicurrency to the blockchain
    Currency:
    - What will the blockchain address look like?
    - What are the formats for the keys necessary to create signatures for transactions?

    Blockchain:
    - One blockchain for both currencies or two?
    - If one there are no pros and mess in UTXO, blocks are bigger, etc


    Key pair(s):
    - Where should the key pair(s) be stored?

    Wallets:
        Should one wallet store key pairs for both currencies
            One address one wallet: differ currencies inside
            Option to differ what currency to work with
            
        Or two wallets are used under the hood so wallet is a logic entity
            One address two wallets: differ wallets
            Differed by the key format


    Wallet:

    UTXO:

### Make UI for wallet