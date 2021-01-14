package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
// Khởi tạo một client hay một user bao gồm:
// + một socket cho client này
// + một send channel để send message
// + một room mà client chat trong đó
type client struct {
	//socket is the web socket for this client
	socket *websocket.Conn

	//send is a channel on which messages are sent
	send chan []byte

	//room is the room this client is chatting in
	room *room
}

// read method cho phép client đọc từ socket thông qua ReadMessage method
// sau đó message đọc được từ socket sẽ được gửi đển forward channel của room type
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}
		c.room.forward <- msg
	}
}

// write method nhận Message từ send channel và
// writing mọi thứ ra ngoài socket bằng WriteMessage method
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
