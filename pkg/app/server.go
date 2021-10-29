package app

import (
	"fmt"
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
	id               int
	timestampSeconds int64
	userName         string
	body             string
}

func NewChatServer() *socketio.Server {
	// CORS: allow all, as we are running locally.
	allowAllOrigins := func(r *http.Request) bool { return true }

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{CheckOrigin: allowAllOrigins},
			&websocket.Transport{CheckOrigin: allowAllOrigins},
		},
	})

	users := map[string]string{}
	pastMessages := []*Message{}

	sendConnectedUsers := func(s socketio.Conn) {
		for _, userName := range users {
			s.Emit("join", userName)
		}
	}

	sendPastMessages := func(s socketio.Conn) {
		for _, msg := range pastMessages {
			s.Emit("msg", msg.id, msg.timestampSeconds, msg.userName, msg.body)
		}
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("--> connected:", s.ID())
		s.Join("chat")
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, userName string) {
		fmt.Println("--> join:", userName)
		sendConnectedUsers(s)
		sendPastMessages(s)
		users[s.ID()] = userName
		server.BroadcastToRoom("/", "chat", "join", userName)
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, body string) {
		messageID++

		id := messageID
		timestampSeconds := time.Now().Unix()
		userName := users[s.ID()]

		msg := &Message{
			id,
			timestampSeconds,
			userName,
			body,
		}
		pastMessages = append(pastMessages, msg)
		fmt.Println("--> msg:", msg)

		server.BroadcastToRoom("/", "chat", "msg", id, timestampSeconds, userName, body)
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		fmt.Println(fmt.Errorf("error: %v", err))
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("--> closed:", reason)
		s.Leave("chat")
		userName := users[s.ID()]
		delete(users, s.ID())
		server.BroadcastToRoom("/", "chat", "leave", userName)
	})

	return server
}
