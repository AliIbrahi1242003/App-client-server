package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync" // Import the sync package for Mutex
)

// This slice will store the entire chat history.
// It's a shared resource, so we need a mutex to protect it.
var chatHistory []string
var mutex = &sync.Mutex{}

// ChatService is the struct that provides the RPC methods.
type ChatService struct{}

// SendMessage is the remote procedure that clients will call.
// It takes a new message (args) and returns the full chat history (reply).
func (s *ChatService) SendMessage(message string, reply *[]string) error {
	// Lock the mutex to prevent other clients from writing at the same time
	mutex.Lock()

	// Add the new message to the shared chat history
	chatHistory = append(chatHistory, message)

	// Set the reply to be the complete, updated chat history
	*reply = chatHistory

	// Unlock the mutex so other clients can now write
	mutex.Unlock()

	fmt.Printf("Received message: %s\n", message)
	return nil
}

func main() {
	// Create a new instance of our ChatService
	chat := new(ChatService)

	// Register the ChatService so it's accessible via RPC
	err := rpc.Register(chat)
	if err != nil {
		log.Fatalf("Error registering RPC service: %s", err)
	}

	// Listen for incoming TCP connections on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error listening on port 1234: %s", err)
	}
	defer listener.Close()

	fmt.Println("Chat Server is running on port 1234...")

	// Infinite loop to accept new connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue // Continue to the next iteration to accept new connections
		}

		// Handle each client connection in a new goroutine (concurrently)
		// This allows the server to handle multiple clients at once.
		go rpc.ServeConn(conn)
	}
}
