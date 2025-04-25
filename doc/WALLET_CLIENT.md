# Wallet Client
A wallet client-server application that works with two cryptocurrencies' coins (referred to as "currencies") and uses [blockchain](https://github.com/pathfinder177/blockchain_go) as the server side.

---

## Wallet (Client Side)

### Features:
1. **Basic Authentication** (password and optionally token)  
   - Where should credentials be stored?
        - Password: DB
        - Token: basically user-side
2. **Create Wallet User Story**
   - User get to page for new user -> register -> wallet is created -> user get to the wallet main page
   - Each pocket has **X value** of currency's coins (e.g., airdrop or distribution)
   
   - Optionally, users can create additional pockets for each of currencies.
        - How to implement it? Probably pocket is a virtual entity as representation of blockchain is unchanged
3. **Delete Wallet**
   - Optionally, user can delete created pockets, but at least one per currency must remain.
   - Otherwise user can delete wallet. Tokens are "burned", wallet is deleted
4. **Send Coins**
   - Incorrect transactions should be reverted to prevent loss of funds
5. **Receive Coins**
6. **Get Data**
   - Balance for currencies
   - Transaction history for each currency over a specific period.
7. ?Wallet smart contract?

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
    - **Users are deleted from DB**
    - **Coins are burned**

## Architecture
The approach is closer to service-oriented architecture at the system level.
And clean architecture at the code level. FIXME(make sure it is done or changed)

Wallet as client-side application implemented as web server
Wallet client uses wallet [server](https://github.com/pathfinder177/blockchain_go) only to connect to blockchain

### Addresses

Wallet server HTTP: localhost:3003
Wallet client HTTP: localhost:3004

### Storage
To implements basic auth: use postgresql as database of users
#### Schema
Users table:
username password

Wallet table:
username wallet

### Handlers:
/
/auth(DB)
/send(bc)
/transactions(all txs history for the wallet for period(7d default))(bc)
/currency_transactions(bc)(txs history for the currency for period(7d default))(bc)
/delete_wallet(DB and blc)

### Wallet UI
- / has:
    - at first stage(without auth):
        - ask for wallet address that created before
        - main page for provided address
- The main page has:
    - address
    - two pockets, one for each currency 
    - buttons for functionality:
        - send
        - get
        - delete wallet

## Projecting
## Features
1. All connections are handled in apart goroutine
    Made part of code async if it makes sense
2. Graceful shutdown
*3. Rate limiting(client)
*4. Advanced routing(Gorilla mux)
5. logs
6. metrics

## Client HTTP Handlers

### User handlers
#### /
Given user goto index
When user submit registration or login data and valid wallet address
Then main page is shown to user

#### /auth(DB) change index page to "main page" is required(mocked)
NEW USER
Given user get to /
When user clicks sign up
Then user is redirected to page with registration form, submit data and redirected to main page
On main page user submits wallet address and it refers to user and can not be reused anyone else

EXISTING USER
Given user get to /
When user click sign in
Then user submit the login data and redirected to main page

##### Client: Clean Architecture
The application is divided into 2 layers, internal and external

Business logic (Go standard library).
Tools (databases, servers, message brokers, any other packages and frameworks).

1) App loads config and runs. Config object uses env vars to store config(12 factors)
- cmd/app/main.go
- config/config.go
- internal:
-   app creates all main objects
        *logger
        repo
        use case
        httpserver/handlers
            run it
        graceful httpserver shutdown
-   controller(server handler layer, http controller) = ENTRYPOINT
-   entity:
        *validation methods
        registration,login entities
-   usecase:
        Repositories, webapi, rpc, and other business logic structures are injected into business logic structures 
        registration,login interfaces
        registration,login usecases
        repo is injected to both usecases
        usecases is created with new
-   repo/persistent: abstract storage to work with
-   repo/webapi: abstract of wallet server

### Server handlers
All handlers to get info from the server side
/get_tx_history
/get_currency_tx_history
/send
/delete*

Make handlers as handleReq function at server side

Go to server part

*use middleware and contextWithValue to register/login and handle user req

How to isolate pages to access login/reg only?

### Clean up the code
### Format the code

## Q&A