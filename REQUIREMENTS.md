# Wallet Client App

A wallet client application that works with two cryptocurrencies (referred to as "currencies") and uses blockchain as the server side.

---

## Key Pair

To send and receive currency, a user needs at least one key pair for each currency.  
Each key pair proves ownership of one wallet address for a specific currency.

### Questions:
- Where should the key pair be stored?

---

## Currencies

### Questions:
- Are cryptocurrencies created in the usual way?
- How can they be integrated into the blockchain?
- Is it possible to mint them directly on the blockchain?
  - If so, they might not need to be created elsewhere.

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
   - Each pocket starts with **X value** of currency (e.g., airdrop or distribution).  
3. **Delete Wallet**  
   - Optionally, users can delete created pockets, but at least one per currency must remain.  
4. **Send Currency**  
   - How to differentiate between currencies when sending?  
   - Incorrect transactions should be reverted to prevent loss of funds.  
5. **Receive Currency**  
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
   - Each pocket contains **X value** of currency.  
5. Get transaction history:  
   - It's empty.  

### **User 2**
6. Start the app.  
7. Register.  
8. Create a wallet.  
9. Get balance of each pocket:  
   - Each pocket contains **X value** of currency.  
10. Get transaction history:  
    - It's empty.  
11. Send **1 unit** of `currency_1` to User 1.  
12. Get balance of pocket with sent currency:  
    - `Balance = previous balance - 1`  
13. Get transaction history:  
    - One transaction (sent to User 1).  

### **User 1**
14. Get balance of pocket with received currency:  
    - `Balance = previous balance + 1`  
15. Get transaction history:  
    - One received transaction.  
16. Send **1 unit** of `currency_2` to User 2.  
17. Get balance of pocket with sent currency:  
    - `Balance = previous balance - 1`  
18. Get transaction history:  
    - One sent transaction.  
    - One received transaction.  

### **User 2**
19. Get balance of pocket with received currency:  
    - `Balance = previous balance + 1`  
20. Get transaction history:  
    - One sent transaction.  
    - One received transaction.  

### **User 1 & User 2**
21. Delete their wallets.  
    - **Coins are gone forever.**  
