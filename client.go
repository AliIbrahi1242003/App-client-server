package main

import (
	"bufio" // To read full lines from standard input
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings" // To trim whitespace and newlines
)

func main() {
	// 1. Connect to the RPC server
	// This handles the requirement: "handle the errors if the server goes down"
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Error dialing server: %s", err)
	}
	defer client.Close()

	fmt.Println("Connected to chat server.")
	fmt.Println("Type your message and press Enter. Type 'exit' to quit.")

	// Create a new reader to read from the keyboard
	// This handles the requirement: "look for something that will read the whole message"
	reader := bufio.NewReader(os.Stdin)

	// 2. Infinite loop to keep the client running
	// This handles the requirement: "client should be running forever"
	for {
		fmt.Print("Enter message: ")

		// Read a full line from the user, until they press Enter (\n)
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}

		// Clean up the input string (remove newline character and extra spaces)
		message = strings.TrimSpace(message)

		// 3. Check for the "exit" command
		if message == "exit" {
			fmt.Println("Disconnecting...")
			break // Break out of the infinite loop
		}

		// 4. Call the remote procedure on the server
		var history []string
		err = client.Call("ChatService.SendMessage", message, &history)
		if err != nil {
			// This handles errors if the server goes down while the client is running
			log.Println("Error calling remote procedure:", err)
			break // Exit the loop if the server connection is lost
		}

		// 5. Print the chat history returned from the server
		fmt.Println("--- Chat History ---")
		for _, msg := range history {
			fmt.Println(msg)
		}
		fmt.Println("----------------------")
	}

	fmt.Println("Goodbye!")
}
