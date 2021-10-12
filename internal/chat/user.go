package chat

import (
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

// User represents a chat room user
type User struct {
	id     int
	conn   *websocket.Conn
	server *Server
	out    chan *Message
	doneCh chan bool
}

const channelBufferSize = 100

var maxID int

// NewUser builds a new user and returns it
func NewUser(conn *websocket.Conn, server *Server) *User {
	if conn == nil {
		panic("Connection cannot be nil")
	}
	if server == nil {
		panic("Server cannot be nil")
	}
	maxID++
	ch := make(chan *Message, channelBufferSize)
	doneCh := make(chan bool)
	return &User{maxID, conn, server, ch, doneCh}
}

func (user *User) Write(message *Message) {
	select {
	case user.out <- message:
	default:
		log.Println(user, "is disconnected")
	}
}

// Listen makes the user listen to message reads and writes in parallel
func (user *User) Listen() {
	go user.listenWrite()
	user.listenRead()
}

func (user *User) listenWrite() {
	for msg := range user.out {
		user.conn.WriteJSON(&msg)
	}
}

func (user *User) listenRead() {
	for {
		select {
		case <-user.doneCh:
			return

		default:
			// Read a message sent by user over websocket
			var message Message
			err := user.conn.ReadJSON(&message)
			if err != nil {
				user.server.Err(err)
				user.server.RemoveUser(user)
				user.doneCh <- true
			} else {
				user.server.AddMessage(&message)
			}
		}
	}
}

func (user *User) String() string {
	return strconv.Itoa(user.id)
}
