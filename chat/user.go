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

// Conn returns the user's websocket connection
func (user *User) Conn() *websocket.Conn {
	return user.conn
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
	select {
	case msg := <-user.out:
		user.conn.WriteJSON(&msg)
	}
}

func (user *User) listenRead() {
	for {
		select {
		case <-user.doneCh:
			user.server.RemoveUser(user)
			return
		default:
			// Read a message sent over the websocket
			var message Message
			err := user.conn.ReadJSON(&message)
			if err != nil {
				user.doneCh <- true
				user.server.Err(err)
			} else {
				user.server.AddMessage(&message)
			}
		}
	}
}

func (user *User) String() string {
	return strconv.Itoa(user.id)
}
