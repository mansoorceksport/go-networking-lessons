
---


# TCP Server-Client Example in Go

> **A progressive journey into TCP networking with Go – from basics to bidirectional chat.**

This repository is a hands-on learning journey into building TCP-based server-client systems using Go. It starts with the basics of TCP communication and progressively introduces more advanced concepts like bidirectional messaging, broadcasting, and client-to-client communication—all through clean, well-structured examples.

Each lesson lives in its own branch (`lesson-1`, `lesson-2`, etc.), so you can follow along step by step and deepen your understanding of real-time networking using raw sockets.


## 📌 Purpose

- Learn how to listen for incoming TCP connections using Go.
- Understand how clients establish a connection and communicate.
- Practice building real-time applications over TCP.
- Incrementally advance your knowledge through practical lessons.

---

## 🗂️ Project Structure

```
.
├── client/
│   └── main.go              # Client implementation
├── server/
│   └── main.go              # Server implementation
├── valueobject/
│   └── constant.go          # Shared constants (e.g., address, port)
```

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go 1.16 or later  
  Check with: `go version`

### 🖥️ Running the Server

```bash
cd server
go run main.go
```

- The server will start listening on the configured port.
- It logs connection events and sends data to connected clients.

### 🧑‍💻 Running the Client

Open a **new terminal**:

```bash
cd client
go run main.go
```

- The client connects to the server and displays any received data.

---

## 🧠 How It Works

- The server continuously sends timestamped messages to each connected client.
- The client connects via TCP and listens for the stream.
- Demonstrates real-time, unidirectional or bidirectional communication over raw TCP.

---

## 🏗️ Lesson Roadmap

| Lesson | Branch        | Description                                 |
|--------|---------------|---------------------------------------------|
| 1      | `lesson-1`    | Basic TCP streaming (server to client)      |
| 2      | `lesson-2`    | Bidirectional server-client messaging       |
| 3      | `lesson-3`    | Server broadcasts messages to all clients   |
| 4      | `lesson-4`    | Client-to-client messaging via server       |
| 5+     | Coming Soon   | Authentication, WebSocket migration, etc.   |

---

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## 🙌 Contributions

This project is meant for learning and growth. Feel free to fork it, improve it, or use it as a foundation for your own networking experiments. Pull requests are welcome!

---

## 🔖 Topics

`golang` • `tcp` • `networking` • `socket-programming` • `client-server` • `real-time` • `learning` • `go-tutorial`
