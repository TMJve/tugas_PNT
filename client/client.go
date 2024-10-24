package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send a message to the server
	fmt.Println("Enter a message to send to the server:")
	message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Fprintf(conn, message)

	// Read response from the server
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Printf("Server response: %s", response)
}
