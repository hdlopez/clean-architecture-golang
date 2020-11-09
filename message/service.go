package message

import (
	"github.com/hdlopez/clean-architecture-golang/apierror"
)

type Service interface {
	Get(ID string) (*Message, error)
}

type messageSrv struct {
	repo Repository
}

func NewService() Service {
	return NewServiceWith(NewRepository())
}

func NewServiceWith(r Repository) Service {
	srv := new(messageSrv)
	srv.repo = r
	return srv
}

func (srv *messageSrv) Get(id string) (*Message, error) {
	// string zero-value is an empty string

	// validate parameters depending your business logic
	if id == "" {
		// bad request
		return nil, apierror.New(400, "Message ID is required")
	}

	// business logic goes here ...

	// retrieve message using its repository
	msg, err := srv.repo.Get(id)

	// do some stuff

	return msg, err
}
