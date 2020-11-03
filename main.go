package main

import (
	"github.com/hdlopez/clean-architecture-golang/api"
)

func main() {
	// Creates an instance of API and calls the Run function
	// to start the application
	api.New().Run()
}
