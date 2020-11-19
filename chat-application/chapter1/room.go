package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	// forward là một channel mà lưu trữ incoming messages
	// mà cần phải được chuyển tiếp đến clients khác
	forward chan []byte

	//join là một channel cho clients muốn được tham gia vào room
	join chan *client
	// leave là một channel cho clients muốn rời khỏi room
	leave chan *client
	// clients lưu trữ toàn bộ clients trong room này
	clients map[*client]bool
}

// khởi tạo room bao gồm các channel và map
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) run() {
	// for có nghĩa là method này sẽ chạy mãi cho đến khi program bị dừng
	for {
		// preceding code sẽ lắng nghe hay watch 3 channel trong room
		// gồm join, leave, forward
		// nếu message nhận được ở bất cứ channel nào thì select statement sẽ
		// chạy code cho trường hợp đó.
		select {
		// nếu nhận được kết quả từ channel join thì thêm client đó vào danh sách clients
		case client := <-r.join:
			r.clients[client] = true
		// nếu nhận được kết quả từ channel leave thì xóa client đó khỏi danh sách
		// và channel client đó sẽ được close
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		// nếu nhận được kết quả từ channle forward thì kết quả được lưu vào msg
		// msg sau đó sẽ được gửi đến toàn bộ client thông qua channel send của client
		// write method của client đó sẽ lấy thông tin đó và gửi nó xuống socket tới browser
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
