package main

import (
	"RedisClone/commands"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
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
		words, err := readCommand(reader)

		if err != nil {
			log.Printf("Client disconnected or error: %v", err)
			return
		}

		result, err := commands.ExecuteCommand(words)
		response := fmt.Sprintf("%s\r\n", result)
		conn.Write([]byte(response))
	}
}

func getCountFromLine(line string) (int, error) {
	count, err := strconv.Atoi(strings.TrimSpace(line[1:]))

	if err != nil {
		return 0, err
	}

	return count, nil
}

func readCommand(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	count, err := getCountFromLine(line)

	if err != nil {
		return nil, err
	}

	words := make([]string, 0, count)

	for range count {
		line, err = reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		length, err := getCountFromLine(line)

		if err != nil {
			return nil, err
		}

		buf := make([]byte, length+2)

		if _, err := io.ReadFull(reader, buf); err != nil {
			return nil, err
		}

		words = append(words, string(buf[:length]))
	}

	return words, nil
}
