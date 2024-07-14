package main

import (
	"DungeonRaceServeur/game"
	"log"
	"net"
	"time"
)

type Server struct {
	startTime             time.Time
	lastCommunicationTime time.Time
	adresse               string
	games                 game.Game
}

func main() {
	server := Server{
		startTime:             time.Now(),
		lastCommunicationTime: time.Now(),
		adresse:               "localhost:8080",
		games:                 game.Game{},
	}

	// Listen for incoming connections
	listener, err := net.Listen("tcp", server.adresse)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer listener.Close()

	log.Printf("Server is listening on %s\n", server.adresse)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		isActive := make(chan bool, 1)
		go handleClient(conn, isActive)
		// <-isActive

		// prevent infinite running
		someoneIsConnected := <-isActive
		log.Println("Is someone connected ?:", someoneIsConnected)
		// select {
		// case mustStop = :
		// 	log.Println("is active:", mustStop)
		// case <-time.After(5 * time.Second):
		// 	log.Println("Timeout: no activity")
		// 	mustStop = true
		// }
		timeOutDuration := 5 * time.Second
		if someoneIsConnected {
			server.lastCommunicationTime = time.Now()
		} else if server.lastCommunicationTime.After(server.lastCommunicationTime.Add(timeOutDuration)) {
			log.Printf("Server stop due to inactivity in the last %s\n", timeOutDuration)
			break
		}
	}
	log.Printf("Server has has successfully ran for: %s and is now stopped\n", time.Since(server.startTime))
}

func handleClient(conn net.Conn, isActive chan bool) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		isActive <- true
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error:", err)
			isActive <- false
			return
		}

		// Process and use the data
		log.Printf("Received: %s\n", buffer[:n])

		// Send back something
		data := []byte("Hello, Client!")
		_, err = conn.Write(data)
		if err != nil {
			log.Println("Error:", err)
			isActive <- false
			return
		}
	}
}
