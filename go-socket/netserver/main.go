package main

import (
	"fmt"
	"net"
	"os"
)

const (
	serverHost = "localhost"
	serverPort = "9999"
	serverType = "tcp"
)

func main() {
	// Create a server liston at the defined address
	svr, err := net.Listen(serverType, serverHost+":"+serverPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// defer is used to ensure that a function call is performed
	// later in a program's execution
	defer svr.Close()
	fmt.Println("Listening on: " + serverHost + ":" + serverPort)
	fmt.Println("Waiting for connection ...")
	for {
		// assign connect from client
		conn, err := svr.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		// each go routine handle one client connect
		go func() {
			handleClient(conn)
		}()
	}
}

// handle client connect
func handleClient(conn net.Conn) {
	for {
		// create a slice of type byte and len =1024
		buff := make([]byte, 1024)
		// connect Read message from client
		msgLen, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		receivedMessage := string(buff[:msgLen])

		//close the connection from client if received the quit message
		if receivedMessage == "quit" {
			clientAddr := conn.RemoteAddr().String()
			fmt.Println("A client: " + clientAddr + " has decided to leave!")
			conn.Close()
			break
		} else {
			fmt.Println("Received: ", receivedMessage)
		}
	}
	return
}
