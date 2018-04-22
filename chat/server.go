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
	incomingMessage := make(chan *Message)
	errorCh := make(chan error)
	doneCh := make(chan bool)
	return &Server{
		connectedUsers,
		messages,
		addUser,
		removeUser,
		incomingMessage,
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
			log.Println("Added user", user, "to chat room")
			server.connectedUsers[user.id] = user
			log.Println(len(server.connectedUsers),
				"are now connected to this chat room.")
			server.sendPastMessages(user)

		case user := <-server.removeUser:
			log.Println("Removing user", user, "from chat room")
			delete(server.connectedUsers, user.id)

		case msg := <-server.newMessage:
			server.Messages = append(server.Messages, msg)
			server.sendAll(msg)

		case err := <-server.errorCh:
			log.Println("Error:", err)

		case <-server.doneCh:
			log.Println("Server stopping...")
			return
		}
	}
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
	conn, _ := upgrader.Upgrade(rw, r, nil)

	var message Message
	err := conn.ReadJSON(message)
	log.Println("Message received:", &message)
	if err != nil {
		log.Println("Error while reading message JSON:", err.Error())
	}

	user := NewUser(conn, server)

	log.Println("Going to add user", user)
	server.AddUser(user)

	log.Println("User added successfully")
	server.AddMessage(&message)
	user.Listen()
}

func (server *Server) handleMessages(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(server)
}
