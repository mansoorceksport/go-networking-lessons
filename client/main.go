package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mansoorceksport/go-networking-lessons/valueobject"
	"log/slog"
	"net"
	"os"
)

var (
	ctx = context.Background()
)

func main() {
	conn, err := net.Dial("tcp", valueobject.ServerAddress)
	if err != nil {
		// error connecting to the server
		slog.ErrorContext(ctx, "failed to connect to the server: %v", slog.Any("err", err.Error()))
		return
	}
	defer conn.Close()

	// Read and print server messages (including asking for client ID)
	incomingScanner := bufio.NewScanner(conn)
	if incomingScanner.Scan() {
		fmt.Println(incomingScanner.Text()) // Enter your client ID:
	}

	// enter client ID
	stdinScanner := bufio.NewScanner(os.Stdin)
	if stdinScanner.Scan() {
		clientId := stdinScanner.Text()
		if clientId == "" {
			slog.ErrorContext(ctx, "client ID cannot be empty")
			return
		}
		_, err = conn.Write([]byte(clientId + "\n"))
		if err != nil {
			slog.ErrorContext(ctx, "failed to send client ID: %v", slog.Any("err", err.Error()))
			return
		}
	}

	// Start a goroutine to handle incoming message
	go func() {
		for incomingScanner.Scan() {
			fmt.Println(incomingScanner.Text())
		}
	}()

	fmt.Println("Type your message in format '@recipient message' (e.g., '@client1 Hello!'):")

	// send messages
	for stdinScanner.Scan() {
		message := stdinScanner.Text()
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			slog.ErrorContext(ctx, "failed to send message: %v", slog.Any("err", err.Error()))
			return
		}

	}
}
