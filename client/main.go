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

func main() {
	ctx := context.Background()
	conn, err := net.Dial("tcp", valueobject.ServerAddress)
	if err != nil {
		// error connecting to the server
		slog.ErrorContext(ctx, "failed to connect to the server: %v", slog.Any("err", err.Error()))
		return
	}
	defer conn.Close()

	slog.InfoContext(ctx, "Connected to streaming server...")
	go handleInput(conn)

	// handle incoming messages from the server
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("server: ", scanner.Text())
		fmt.Print("> ")
	}

	if err := scanner.Err(); err != nil {
		slog.ErrorContext(ctx, "error reading from the server: %v", slog.Any("err", err.Error()))
		os.Exit(1)
	}

}

func handleInput(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte("client: " + message + "\n"))
		if err != nil {
			fmt.Println("Error sending message to server:", err)
			return
		}

		// Print a new prompt line after sending the message
		fmt.Print("> ")

		// exit if the message is "exit"
		if message == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
			return
		}
	}
}
