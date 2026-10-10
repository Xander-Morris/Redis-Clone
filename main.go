package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6380")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() 

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("Client disconnected or error: %v", err)
			return 
		}

		_, err = conn.Write([]byte("ACK: " + message))
		
		if err != nil {
			log.Printf("Failed to write to client: %v", err)
			return
		}
	}
}