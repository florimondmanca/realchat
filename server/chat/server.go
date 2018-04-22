package chat

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Server listens to events
type Server struct {
	connectedUsers map[int]*User
	Messages       []*Message `json:"messages"`
	addUser        chan *User
	removeUser     chan *User
	newMessage     chan *Message
	errorCh        chan error
	doneCh         chan bool
}

// NewServer builds a server and returns it
func NewServer() *Server {
	connectedUsers := make(map[int]*User)
	messages := []*Message{}
	addUser := make(chan *User)
	removeUser := make(chan *User)
	newMessage := make(chan *Message)
	errorCh := make(chan error)
	doneCh := make(chan bool)
	return &Server{
		connectedUsers,
		messages,
		addUser,
		removeUser,
		newMessage,
		errorCh,
		doneCh,
	}
}

// Listen to events on the server
func (server *Server) Listen() {
	log.Println("Server listening...")

	http.HandleFunc("/chat", server.handleChat)
	http.HandleFunc("/messages", server.handleMessages)

	for {
		select {
		case user := <-server.addUser:
			log.Println("[ADD USER]", user)
			server.connectedUsers[user.id] = user
			server.logConnected()
			server.sendPastMessages(user)

		case user := <-server.removeUser:
			log.Println("[REMOVE USER]", user)
			delete(server.connectedUsers, user.id)
			server.logConnected()

		case msg := <-server.newMessage:
			log.Println("[MESSAGE]", msg)
			server.Messages = append(server.Messages, msg)
			server.sendAll(msg)

		case err := <-server.errorCh:
			log.Println("[ERROR]", err.Error())

		case <-server.doneCh:
			log.Println("Server stopping...")
			return
		}
	}
}

func (server *Server) logConnected() {
	log.Println(len(server.connectedUsers), "users now connected")
}

// Err seends an error to the server
func (server *Server) Err(err error) {
	server.errorCh <- err
}

func (server *Server) sendPastMessages(user *User) {
	for _, msg := range server.Messages {
		user.Write(msg)
	}
}

func (server *Server) sendAll(msg *Message) {
	for _, user := range server.connectedUsers {
		user.Write(msg)
	}
}

// AddUser registers a user for addition to list of connected users
func (server *Server) AddUser(user *User) {
	server.addUser <- user
}

// RemoveUser registers a user for removal to list of connected users
func (server *Server) RemoveUser(user *User) {
	server.removeUser <- user
}

// AddMessage registers a new incoming message on the server
func (server *Server) AddMessage(msg *Message) {
	server.newMessage <- msg
}

func (server *Server) handleChat(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		server.Err(err)
		return
	}
	user := NewUser(conn, server)
	server.AddUser(user)
	user.Listen()
}

func (server *Server) handleMessages(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(server)
}
