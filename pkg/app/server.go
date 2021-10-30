package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var messageID = 0

type Message struct {
	Id               int    `json:"id"`
	TimestampSeconds int64  `json:"timestampSeconds"`
	UserName         string `json:"userName"`
	Body             string `json:"body"`
}

type Clock interface {
	Now() int64
}

type DefaultClock struct {
	Clock
}

func (c *DefaultClock) Now() int64 {
	return time.Now().Unix()
}

type Server struct {
	sio          *socketio.Server
	pastMessages []*Message
	users        map[string]string
	Clock        Clock
}

var (
	defaultClock = &DefaultClock{}
)

func NewChatServer() *Server {
	server := &Server{
		sio:          nil,
		pastMessages: []*Message{},
		users:        map[string]string{},
		Clock:        defaultClock,
	}
	server.sio = server.createSocketIO()
	return server
}

func (s *Server) Handler() http.Handler {
	return s.sio
}

func (s *Server) Start() {
	go s.sio.Serve()
}

func (s *Server) Stop() {
	s.sio.Close()
}

func (s *Server) createSocketIO() *socketio.Server {
	// CORS: allow all, as we are running locally.
	allowAllOrigins := func(r *http.Request) bool { return true }

	sio := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{CheckOrigin: allowAllOrigins},
			&websocket.Transport{CheckOrigin: allowAllOrigins},
		},
	})

	sio.OnConnect("/", func(conn socketio.Conn) error {
		log.Println("--> connected:", conn.ID())
		conn.Join("chat")
		return nil
	})

	sio.OnEvent("/", "join", func(conn socketio.Conn, userName string, channel string) {
		log.Println("--> join:", userName, channel)
		conn.Join(getRoom(channel))
		s.sendHistory(conn)
		s.addUser(conn.ID(), userName)
		sio.BroadcastToRoom("/", "chat", "join", userName)
	})

	sio.OnEvent("/", "msg", func(conn socketio.Conn, body string) {
		msg := s.makeMessage(conn.ID(), body)
		log.Println("--> msg:", msg)
		s.pastMessages = append(s.pastMessages, msg)
		sio.BroadcastToRoom("/", "chat", "msg", msg.Id, msg.TimestampSeconds, msg.UserName, msg.Body)
	})

	sio.OnError("/", func(_ socketio.Conn, err error) {
		log.Println(fmt.Errorf("error: %v", err))
	})

	sio.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		conn.Leave("chat")
		userName := s.popUser(conn.ID())
		log.Println("--> leave:", userName)
		sio.BroadcastToRoom("/", "chat", "leave", userName)
	})

	return sio
}

func (s *Server) addUser(id string, userName string) {
	s.users[id] = userName
}

func (s *Server) getUser(id string) string {
	return s.users[id]
}

func (s *Server) popUser(id string) string {
	userName := s.users[id]
	delete(s.users, id)
	return userName
}

func (s *Server) makeMessage(userId string, body string) *Message {
	messageID++
	id := messageID
	timestampSeconds := s.Clock.Now()
	userName := s.getUser(userId)
	return &Message{id, timestampSeconds, userName, body}
}

func (s *Server) sendHistory(conn socketio.Conn) {
	for _, userName := range s.users {
		conn.Emit("join", userName)
	}

	for _, msg := range s.pastMessages {
		conn.Emit("msg", msg.Id, msg.TimestampSeconds, msg.UserName, msg.Body)
	}
}

func getRoom(channel string) string {
	return fmt.Sprintf("channel.%s", channel)
}
