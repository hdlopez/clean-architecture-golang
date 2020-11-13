package main

import (
	"github.com/hdlopez/clean-architecture-golang/api"
	"github.com/hdlopez/clean-architecture-golang/config"
	"github.com/hdlopez/clean-architecture-golang/message"
	"github.com/hdlopez/clean-architecture-golang/restclient"
)

func main() {
	// Dependency injection section
	clients := config.RestClients()

	msgAPI := restclient.NewMessageAPI(clients[config.MessageAPI])

	msgRepo := message.NewRepository(msgAPI)

	msgSrv := message.Service(msgRepo)
	msgCtrl := api.NewMessageController(msgSrv)

	// Creates the API intance
	api := api.New(msgCtrl)

	// Runs the application
	api.Run()
}
