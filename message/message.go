package message

import "github.com/hdlopez/clean-architecture-golang/restclient"

type Message struct {
	Text string `json:"text"`
}

func build(msg *restclient.Message) *Message {
	return &Message{msg.Text}
}
