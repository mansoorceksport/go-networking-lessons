# TCP Server-Client Example

This project demonstrates a simple TCP server-client setup in Go.

## Purpose

- Show how to listen for incoming TCP connections.
- Establish a client connection to receive streamed data.
- Illustrate how data is sent from the server and received by the client.

## Project Structure

- `server/main.go`: Implements the server that accepts connections and sends data.
- `client/main.go`: Implements the client that connects to the server and receives data.
- `valueobject/constant.go`: Contains shared constants (like the server address).

## How to Use

### Prerequisites
- Go 1.16 or later.

### Running the Server
1. Navigate to `server`:  
   \`cd server\`
2. Run:  
   \`go run main.go\`
3. The server will listen on the configured port and log connection events.

### Running the Client
1. Open a new terminal and navigate to `client`:  
   \`cd client\`
2. Run:  
   \`go run main.go\`
3. The client will connect and display any data received from the server.

## Explanation
- The server waits for clients to connect and periodically streams data.
- Each client reads the incoming data and prints it to its console.