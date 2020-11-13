package config

import (
	"github.com/go-resty/resty/v2"
)

const (
	MessageAPI = "MessageAPI"
)

func RestClients() map[string]*resty.Client {
	restClients := map[string]*resty.Client{
		MessageAPI: resty.New(),
		// Here we can add rest clients for your APIs.

		// You must probably need to configure timeouts
		// and retries strategies for all your rest clients.
	}

	return restClients
}
