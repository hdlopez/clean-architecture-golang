package message

import "github.com/hdlopez/clean-architecture-golang/restclient"

type Repository interface {
	Get(id string) (*Message, error)
}

type msgRepo struct {
	// It can be easly changed by a database or other storage without touching
	// business logic
	api restclient.MessageAPI
}

func NewRepository() Repository {
	repo := new(msgRepo)
	repo.api = restclient.NewMessageAPI()
	return repo
}

func (repo *msgRepo) Get(id string) (*Message, error) {
	// Getting message from the external API.
	msg, err := repo.api.Get(id)
	if err != nil {
		return nil, err
	}

	return build(msg), nil
}
