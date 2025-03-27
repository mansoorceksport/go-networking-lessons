package main

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"tutorials/tcp/valueobject"
)

func main() {
	ctx := context.Background()
	conn, err := net.Dial("tcp", valueobject.ServerAddress)
	if err != nil {
		// error connecting to the server
		slog.ErrorContext(ctx, "failed to connect to the server: %v", slog.Any("err", err.Error()))
	}
	defer conn.Close()

	slog.InfoContext(ctx, "Connected to streaming server...")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		slog.InfoContext(ctx, fmt.Sprintf("Received line: %s", scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		slog.ErrorContext(ctx, "error reading from the server: %v", slog.Any("err", err.Error()))
		os.Exit(1)
	}
}
