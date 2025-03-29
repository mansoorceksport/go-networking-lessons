

```markdown
# Lesson 4: Client-to-Client Communication via Server

This lesson demonstrates how clients can send messages to each other through a central TCP server. The server handles routing and forwarding messages between connected clients based on their unique IDs.

---

## 🧠 What You'll Learn

- How to assign and use client IDs.
- How to route messages from one client to another through the server.
- Parsing message formats to determine the target recipient.
- Managing multiple client connections with Go's concurrency model.
```

---

## 🗂️ Project Structure

```
.
├── client/
│   └── main.go              # Prompts for client ID and handles messaging
├── server/
│   └── main.go              # Registers client IDs and forwards messages
├── valueobject/
│   └── constant.go          # Shared constants (like server address)
```

---

## 🛠️ How It Works

1. Each client is asked to enter a unique **Client ID** when connecting.
2. The server stores all connected clients in memory.
3. Clients can send messages using the format:

```
@recipient-id Your message here
``` 

4. The server reads the sender's message, finds the recipient, and forwards it.

---

## 🚀 Getting Started

### 1. Run the Server

```bash
cd server
go run main.go
```

You’ll see logs when clients connect or disconnect.

---

### 2. Run the Clients (in separate terminals)

```bash
cd client
go run main.go
```

You’ll be prompted to enter a **unique client ID**. After that, you can send messages in this format:

```
@client2 Hello, how are you?
```

If `client2` is connected, they will receive:

```
From client1: Hello, how are you?
```

---

## 📝 Example Interaction

**Client 1 terminal:**
```
Enter your client ID: client1
@client2 Hello from client1!
```

**Client 2 terminal:**
```
Enter your client ID: client2
From client1: Hello from client1!
```

---

## 📌 Notes

- Clients must be connected and properly registered with unique IDs.
- Messages that don’t match the `@recipient` format are ignored (or could be improved with validation).
- You can enhance this lesson further by supporting:
  - `@all` broadcasting
  - Command parsing (e.g., `/list` to list online users)
  - Authentication

---

## 🏁 Next Steps

This lesson sets the foundation for more advanced features like private/public channels, user lists, and real-time WebSocket-based communication.

Check out the other lessons in the `go-networking-lessons` repo to continue learning.

---