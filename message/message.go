package message

import "github.com/hdlopez/clean-architecture-golang/restclient"

// Message entity
type Message struct {
	Text string `json:"text"`
}

// build creates a new Message from restclient.Message
func build(msg *restclient.Message) *Message {
	return &Message{msg.Text}
}
