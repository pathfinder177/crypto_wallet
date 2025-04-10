# Wallet Client App
'
A wallet client application that works with two cryptocurrencies' coins (referred to as "currencies") and uses blockchain as the server side.
Thus blockchain should be multicurrency(WIP: https://github.com/pathfinder177/blockchain_go)

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

### Make wallet multicurrency


### Wallet UI

### Make UI for wallet
