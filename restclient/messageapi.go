package restclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var (
	// messageAPIURL is the an example API url to get messages
	messageAPIURL = "http://localhost:3000/messages/%s"
)

// MessageAPI interface defines the methods for message API client
type MessageAPI interface {
	Get(ctx context.Context, id string) (*Message, error)
}

type msgAPI struct {
	restAPI
}

// Message entity for restclient
type Message struct {
	Text string `json:"text"`
}

// NewMessageAPI creates a new instance of MessageAPI
func NewMessageAPI(c *resty.Client) MessageAPI {
	api := new(msgAPI)
	api.readClient = c
	return api
}

func (api *msgAPI) readURL(id string) string {
	return fmt.Sprintf(messageAPIURL, id)
}

// Get a message from our external Message API
func (api *msgAPI) Get(ctx context.Context, id string) (*Message, error) {
	url := api.readURL(id)

	msg := new(Message)
	err := api.get(ctx, url, http.Header{}, msg)

	return msg, err
}
