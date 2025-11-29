package config

import (
	"github.com/go-resty/resty/v2"
)

const (
	// MessageAPI is the name of the message API client
	MessageAPI = "MessageAPI"
)

// RestClients returns a map of configured resty clients
func RestClients() map[string]*resty.Client {
	restClients := map[string]*resty.Client{
		MessageAPI: resty.New(),
		// Here we can add rest clients for your APIs.

		// You must probably need to configure timeouts
		// and retries strategies for all your rest clients.
	}

	return restClients
}
