package message

import (
	"context"

	"github.com/hdlopez/clean-architecture-golang/apierror"
)

// Service interface defines the methods for message service
type Service interface {
	Get(ctx context.Context, ID string) (*Message, error)
}

type messageSrv struct {
	repo Repository
}

// NewService creates a new instance of Service
func NewService(r Repository) Service {
	srv := new(messageSrv)
	srv.repo = r
	return srv
}

func (srv *messageSrv) Get(ctx context.Context, id string) (*Message, error) {
	// string zero-value is an empty string

	// validate parameters depending your business logic
	if id == "" {
		// bad request
		return nil, apierror.New(400, "Message ID is required")
	}

	// business logic goes here ...

	// retrieve message using its repository
	msg, err := srv.repo.Get(ctx, id)

	// do some stuff

	return msg, err
}
