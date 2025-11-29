package message

import (
	"context"

	"github.com/hdlopez/clean-architecture-golang/restclient"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Message, error)
}

type msgRepo struct {
	// It can be easly changed by a database or other storage without touching
	// business logic
	api restclient.MessageAPI
}

func NewRepository(api restclient.MessageAPI) Repository {
	repo := new(msgRepo)
	repo.api = api
	return repo
}

func (repo *msgRepo) Get(ctx context.Context, id string) (*Message, error) {
	// Getting message from the external API.
	msg, err := repo.api.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return build(msg), nil
}
