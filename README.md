# Backend API for Merchant & Bank Payment

This project is a backend API designed to simulate interactions between registered account merchants and customers in banks using JSON files including activity logging. It features token-based authentication (JWT).

## Features
- **Login**: login for customer (create access token)
- **Payments**: Customer-to-Merchant
- **Logout**: Ends user session (blacklist access token).
- **Activity Logging**: All actions are logged in a history file.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/alwafie/bank-payment-api
   cd bank-payment-api
   ```
2. Install dependencies:
    ```bash
   go mod init
   go get github.com/gorilla/mux github.com/golang-jwt/jwt/v5
   ```
3. Run the application:
    ```bash
   go run main.go
   ```

## API Endpoints
1. **Customer Login**
    
    POST `/auth/login` 

      Request
   ```json
   {
      "username" : "joni",
      "password" : "password123"
   }
   ```
    Response
    ```json
   {
      "customerId" : "c1"
      "token" : "<JWT_TOKEN>"
   }
   ```
2. **Payments**

   POST `/make/payment`

   Request
   ```json
   {
      "merchantAccountNumber" : "12344321",
      "amount" : 100000
   }
   ```
   Response
    ```json
   {
      "amount": 100000,
      "customerId": "c1",
      "customerName": "Joni",
      "merchantAccountNumber": "12344321",
      "merchantName": "Tokonyadia",
      "status": "success"
    }
   ```
3. **Customer Logout**

   POST `/auth/logout`
   Response
    ```json
    {
      "message": "Logout successful"
    }

## JSON Data Files
- `customers.json` = Customer data
- `merchants.json` = Merchant data
- `history.json` = Logs all action
- `blacklist.json` = Blacklisted token

## Example Logs
```json
  {"action":"login","customerId":"c1","customerName":"Joni","timestamp":"2024-11-25T20:32:47.211989032+07:00"}
  {"action":"payment","customerId":"c1","customerName":"Joni","customerAccountNumber":"123456789","merchantName":"Tokonyadia","merchantNumberAccount":"12344321","timestamp":"2024-11-25T20:33:15.111944263+07:00","status":"payment"}
  {"action":"logout","customerId":"c1","customerName":"Joni","timestamp":"2024-11-25T20:34:15.065654751+07:00"}

```
