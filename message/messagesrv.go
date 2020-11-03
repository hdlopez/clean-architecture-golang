package message

import (
	"github.com/hdlopez/clean-architecture-golang/apierror"
	"github.com/hdlopez/clean-architecture-golang/restclient"
)

type MessageSrv struct {
	api messageAPI
}

type messageAPI interface {
	Get(id string) (*restclient.Message, error)
}

func NewService() *MessageSrv {
	srv := new(MessageSrv)
	srv.api = restclient.NewMessageAPI()
	return srv
}

func (srv *MessageSrv) Get(id string) (*Message, error) {
	// string zero-value is an empty string
	if id == "" {
		// bad request
		return nil, apierror.New(400, "Message ID is required")
	}

	return srv.Get(id)
}
