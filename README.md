# 🚀 Simple RPC Chatroom (Distributed Systems Assignment)

This project is a simple, terminal-based chatroom application built in Go. It was developed as a solution for *Assignment 04* of the *Distributed Systems (Level 4)* course.

The application uses Go's built-in net/rpc package to demonstrate a basic client-server architecture. Multiple clients can connect to a central server, send messages, and receive the entire updated chat history with each message sent.

---

## 🎥 Project Demonstration

As required by the assignment instructions, here is a screen recording of the application running. It shows the server being initialized, followed by multiple clients connecting, sending messages, and receiving the synchronized chat history.

[Project Demonstration Video](https://drive.google.com/file/d/12vUr7sMlVuJQ2wqCATiUwsoI6O2W3Qv9/view?usp=drivesdk)

---

## ✨ Features

* *Centralized Server:* A single server manages and stores all messages in a shared chat history.
* *Multi-Client Support:* The server uses goroutines (go rpc.ServeConn(conn)) to handle multiple client connections concurrently.
* *Thread-Safe:* Uses a sync.Mutex on the server to prevent race conditions when multiple clients write to the chat history at the same time.
* *Full-Line Input:* The client correctly reads full-line messages (including spaces) using bufio.NewReader.
* *Graceful Exit:* Clients can disconnect cleanly by typing exit.
* *Error Handling:* The client includes basic error handling for server connection issues (both on startup and during a call).

---

## 🛠 Technology Stack

* *Language:* *Go (Golang)*
* *Core Packages:*
    * net/rpc: For the Remote Procedure Call mechanism.
    * net: For TCP network listening (net.Listen) and dialing (rpc.Dial).
    * sync: For Mutex to ensure thread safety on the shared chatHistory.
    * bufio: For reading full-line user input from os.Stdin.

---

## 🚀 How to Run

To run this project, you will need to have [Go](https://golang.org/dl/) installed on your machine.

### 1. Clone the Repository

```bash
git clone [https://github.com/AliIbrahi1242003/App-client-server.git](https://github.com/AliIbrahi1242003/App-client-server.git)
cd App-client-server

2. Run the Server
Open a terminal and run the following command to start the server:
go run server.go

You should see the output: Chat Server is running on port 1234...
3. Run the Client
Open a new terminal (or multiple new terminals to simulate multiple users):
go run client.go

You should see the output: Connected to chat server.
4. Start Chatting!
 * Type any message in a client terminal and press Enter.
 * The message will be sent to the server, and the full chat history (including your new message) will be printed.
 * Go to another client terminal and send a message. You will see the history from all clients.
 * To disconnect a client, type exit and press Enter.
```

---

## Installation Instructions

1. Ensure you have Go installed on your machine. You can download it from [Go Downloads](https://golang.org/dl/).
2. Clone the repository:
   ```bash
   git clone https://github.com/AliIbrahi1242003/App-client-server.git
   cd App-client-server
   ```
3. Install any necessary dependencies (if applicable).