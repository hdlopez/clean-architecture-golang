package restclient

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	// userAPIURL is a context url to config API
	messageAPIURL = "http://localhost:3000/messages/%s"
)

type MessageAPI interface {
	Get(id string) (*Message, error)
}

type msgAPI struct {
	restAPI
}

type Message struct {
	Text string `json:"text"`
}

func NewMessageAPI() MessageAPI {
	api := new(msgAPI)

	// Create a Resty Client
	client := resty.New()
	// You probably need to configure timeouts
	// and retries strategies for message API

	api.readClient = client

	//api.logger = &apiLogger{}
	return api
}

func (api *msgAPI) readURL(id string) string {
	return fmt.Sprintf(messageAPIURL, id)
}

func (api *msgAPI) Get(id string) (*Message, error) {
	url := api.readURL(id)

	msg := new(Message)
	res, err := api.get(url, http.Header{}, msg)

	// type assertion
	msg, _ = res.(*Message)

	return msg, err
}
