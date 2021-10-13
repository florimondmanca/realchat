package chat

import (
	"fmt"
	"time"
)

var maxMessageId int

type Type string

const (
	CHAT  Type = "CHAT"
	JOIN  Type = "JOIN"
	LEAVE Type = "LEAVE"
)

// Message stores information about a chat message
type Message struct {
	Id               int                    `json:"id"`
	Type             Type                   `json:"type"`
	Data             map[string]interface{} `json:"data"`
	TimestampSeconds int64                  `json:"timestampSeconds"`
}

// ChatMessageData defines the expected shape of user message payloads
type ChatMessageData struct {
	UserName string `json:"userName"`
	Body     string `json:"body"`
}

func newMessage() *Message {
	msg := &Message{Id: maxMessageId, TimestampSeconds: time.Now().Unix()}
	maxMessageId++
	return msg
}

func NewChatMessage(userName string, body string) *Message {
	msg := newMessage()
	msg.Type = CHAT
	data := make(map[string]interface{})
	data["userName"] = userName
	data["body"] = body
	msg.Data = data
	return msg
}

func NewJoinMessage(userName string) *Message {
	msg := newMessage()
	msg.Type = JOIN
	data := make(map[string]interface{})
	data["userName"] = userName
	msg.Data = data
	return msg
}

func NewLeaveMessage(userName string) *Message {
	msg := newMessage()
	msg.Type = LEAVE
	data := make(map[string]interface{})
	data["userName"] = userName
	msg.Data = data
	return msg
}

func (message *Message) String() string {
	return fmt.Sprintf("[%d] %s: %v",
		message.TimestampSeconds, message.Type, message.Data)
}
