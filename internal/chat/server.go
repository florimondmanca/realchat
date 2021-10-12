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

// NewServer builds a chat server and returns it
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

// HandleChat handles new connections
func (s *Server) HandleChat(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		s.Err(err)
		return
	}
	user := NewUser(conn, s)
	s.AddUser(user)
	user.Listen()
}

// HandleMessages handles chat messages
func (s *Server) HandleMessages(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(s)
}

// Listen to events on the server
func (s *Server) Listen() {
	log.Println("Server listening...")

	for {
		select {
		case user := <-s.addUser:
			log.Println("[ADD USER]", user)
			s.connectedUsers[user.id] = user
			s.logConnected()
			s.sendPastMessages(user)

		case user := <-s.removeUser:
			log.Println("[REMOVE USER]", user)
			delete(s.connectedUsers, user.id)
			s.logConnected()

		case msg := <-s.newMessage:
			log.Println("[MESSAGE]", msg)
			s.Messages = append(s.Messages, msg)
			s.notifyAll(msg)

		case err := <-s.errorCh:
			log.Println("[ERROR]", err.Error())

		case <-s.doneCh:
			return
		}
	}
}

// Stop stops the chat server
func (s *Server) Stop() {
	close(s.doneCh)
}

func (s *Server) logConnected() {
	log.Println(len(s.connectedUsers), "users now connected")
}

// Err seends an error to the server
func (s *Server) Err(err error) {
	s.errorCh <- err
}

func (s *Server) sendPastMessages(user *User) {
	for _, msg := range s.Messages {
		user.Write(msg)
	}
}

func (s *Server) notifyAll(msg *Message) {
	for _, user := range s.connectedUsers {
		user.Write(msg)
	}
}

// AddUser registers a user for addition to list of connected users
func (s *Server) AddUser(user *User) {
	s.addUser <- user
}

// RemoveUser registers a user for removal to list of connected users
func (server *Server) RemoveUser(user *User) {
	server.removeUser <- user
}

// AddMessage registers a new incoming message on the server
func (server *Server) AddMessage(msg *Message) {
	server.newMessage <- msg
}
