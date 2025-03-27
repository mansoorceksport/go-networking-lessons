package main

import (
	"context"
	"fmt"
	"github.com/mansoorceksport/go-networking-lessons/valueobject"
	"log/slog"
	"net"
	"time"
)

func main() {
	ctx := context.Background()
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
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		message := fmt.Sprintf("Streaming data: %v\n", time.Now().Format(time.RFC3339))

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}

		time.Sleep(1 * time.Second)
	}

}
