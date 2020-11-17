package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

const (
	serverHost = "localhost"
	serverPort = "9999"
	serverType = "tcp"
)

func main() {
	fmt.Println("client is running")
	// connect to server
	conn, err := net.Dial(serverType, serverHost+":"+serverPort)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// loop to send message to the server
	// if quit close the connection
	for {
		message := getMessage()
		_, err = conn.Write([]byte(message))
		if message == "quit" {
			fmt.Println("Good bye!")
			break
		}
	}
}

// get message from keyboard
func getMessage() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	mess, _ := reader.ReadString('\n')

	// delete \n in the input string
	if runtime.GOOS == "windows" {
		mess = strings.TrimRight(mess, "\r\n")
	} else {
		mess = strings.TrimRight(mess, "\n")
	}
	return mess
}
