package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

// API type
type API struct {
	pingCtrl *pingCtrl
	msgCtrl  MessageController
}

// New creates a new instance of API type
func New(msgCtrl MessageController) *API {
	api := new(API)
	api.msgCtrl = msgCtrl
	api.pingCtrl = newPingCtrl()
	return api
}

// Ctx interface is an abstraction of gin.Context.
// This interface makes mocking on unit testing easier
type Ctx interface {
	context.Context
	Param(key string) string

	String(code int, format string, values ...interface{})
	JSON(code int, obj interface{})
}

// Run function runs our API application
func (api *API) Run() {
	r := gin.Default()

	api.configRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// configRoutes set all the API endpoints
func (api *API) configRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) { api.pingCtrl.Ping(c) })
	r.GET("/messages/:id", func(c *gin.Context) { api.msgCtrl.Get(c) })
}
