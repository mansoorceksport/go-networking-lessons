package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mansoorceksport/go-networking-lessons/valueobject"
	"log/slog"
	"net"
	"strings"
	"sync"
)

var (
	clients   = make(map[string]net.Conn)
	clientsMu = &sync.Mutex{}
	ctx       = context.Background()
)

func main() {
	listener, err := net.Listen("tcp", valueobject.ServerAddress)
	if err != nil {
		slog.ErrorContext(ctx, "failed to listen: %v", slog.Any("err", err))
	}
	defer listener.Close()

	slog.InfoContext(ctx, "server is listening on :9000")

	for {
		// accept the connection
		conn, err := listener.Accept()
		if err != nil {
			slog.ErrorContext(ctx, "failed to accept connection:", slog.Any("err", err))
			continue
		}

		slog.InfoContext(ctx, "accepted new connection", slog.Any("IP", conn.RemoteAddr().String()))
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer func() {
		conn.Close()
		slog.Info("Client disconnected:", slog.Any("IP", conn.RemoteAddr().String()))
	}()

	// request client id
	conn.Write([]byte("Enter your client ID: " + "\n"))
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return
	}
	clientId := strings.TrimSpace(scanner.Text())
	if clientId == "" {
		conn.Write([]byte("Invalid client ID. Disconnecting...\n"))
		return
	}

	clientsMu.Lock()
	clients[clientId] = conn
	clientsMu.Unlock()

	fmt.Println("client connected:", clientId)
	conn.Write([]byte("Welcome " + clientId + "! You can start sending messages.\n"))

	// Handle incoming messages
	for scanner.Scan() {
		message := scanner.Text()
		handleMessage(clientId, message)

	}

	if err := scanner.Err(); err != nil {
		slog.ErrorContext(context.Background(), "error reading from client:", slog.Any("err", err))
	}

	// Cleanup on client disconnect
	clientsMu.Lock()
	delete(clients, clientId)
	clientsMu.Unlock()
	fmt.Println("Client disconnected:", clientId)
}

func handleMessage(senderId, message string) {
	// Message format: @clientId message
	// Example: @client1 Hello
	if !strings.HasPrefix(message, "@") {
		return
	}
	// Split the message into parts
	parts := strings.SplitN(message[1:], " ", 2)
	if len(parts) < 2 {
		return
	}
	recipientId, message := parts[0], parts[1]
	clientsMu.Lock()
	recipientConn, ok := clients[recipientId]
	clientsMu.Unlock()
	if !ok {
		fmt.Println("Client not found:", recipientId)
		return
	}

	// Send the message to the recipient
	_, err := recipientConn.Write([]byte(senderId + ": " + message + "\n"))
	if err != nil {
		fmt.Println("Error sending message to recipient:", err)
		clientsMu.Lock()
		delete(clients, recipientId) // remove the disconnected client
		clientsMu.Unlock()
	}
}
