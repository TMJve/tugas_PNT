package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Start a server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read message from the client
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}
	fmt.Printf("Received message: %s", message)

	// Send a response to the client
	response := "Hello from the server!\n"
	conn.Write([]byte(response))
}
