package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mansoorceksport/go-networking-lessons/valueobject"
	"log/slog"
	"net"
	"os"
	"sync"
)

var (
	clients   = make(map[net.Conn]bool)
	clientsMu = &sync.Mutex{}
)

func main() {
	ctx := context.Background()
	listener, err := net.Listen("tcp", valueobject.ServerAddress)
	if err != nil {
		slog.ErrorContext(ctx, "failed to listen: %v", slog.Any("err", err))
	}
	defer listener.Close()

	slog.InfoContext(ctx, "server is listening on :9000")

	// Goroutine to handle sending messages from server terminal
	go handleServerInputs()
	fmt.Print("> you: ")

	for {
		// accept the connection
		conn, err := listener.Accept()
		if err != nil {
			slog.ErrorContext(ctx, "failed to accept connection:", slog.Any("err", err))
			continue
		}

		clientsMu.Lock()
		clients[conn] = true
		clientsMu.Unlock()

		slog.InfoContext(ctx, "accepted new connection", slog.Any("IP", conn.RemoteAddr().String()))
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		slog.Info("Client disconnected:", slog.Any("IP", conn.RemoteAddr().String()))
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
		fmt.Print("> ")

	}

	if err := scanner.Err(); err != nil {
		slog.ErrorContext(context.Background(), "error reading from client:", slog.Any("err", err))
	}

}

func handleServerInputs() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "" {
			return
		}
		// Print a new prompt line after sending the message
		fmt.Print("> ")

		broadcastMessage(message)

	}
}

func broadcastMessage(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for client := range clients {
		_, err := client.Write([]byte(message + "\n"))
		if err != nil {
			slog.Error("error writing to client:", slog.Any("err", err))
			client.Close()
			delete(clients, client) // remove the disconnected client
		}
	}
}
