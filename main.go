package main

import (
	"github.com/hdlopez/clean-architecture-golang/api"
	"github.com/hdlopez/clean-architecture-golang/config"
	"github.com/hdlopez/clean-architecture-golang/message"
	"github.com/hdlopez/clean-architecture-golang/restclient"
)

func main() {
	// Create a Resty Client
	clients := config.RestClients()

	msgAPI := restclient.NewMessageAPI(clients[config.MessageAPI])

	msgRepo := message.NewRepository(msgAPI)

	msgSrv := message.Service(msgRepo)
	msgCtrl := api.NewMessageController(msgSrv)
	// Creates an instance of API
	api := api.New(msgCtrl)

	// calls the Run function
	// to start the application
	api.Run()
}
