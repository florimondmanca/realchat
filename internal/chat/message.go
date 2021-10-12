package chat

import "fmt"

// Message stores information about a chat message
type Message struct {
	Id               int    `json:"id"`
	UserName         string `json:"userName"`
	Body             string `json:"body"`
	Action           string `json:"action"`
	TimestampSeconds int64  `json:"timestampSeconds"`
}

func (message *Message) String() string {
	return fmt.Sprintf("[%d] %s %s: %s",
		message.TimestampSeconds, message.UserName, message.Action, message.Body)
}
