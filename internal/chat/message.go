package chat

import "fmt"

// Message stores information about a chat message
type Message struct {
	UserName  string `json:"userName"`
	Body      string `json:"body"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}

func (message *Message) String() string {
	return fmt.Sprintf("[%s] %s %s: %s",
		message.Timestamp, message.UserName, message.Action, message.Body)
}
