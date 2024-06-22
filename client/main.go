package main

import (
	"log"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// Read and process data from the server
	// Create a buffer to read data into
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Printf("Received: %s\n", buffer[:n])
}
