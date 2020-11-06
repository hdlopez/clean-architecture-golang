package api

import (
	"github.com/gin-gonic/gin"
)

// API type
type API struct {
	pingCtrl *pingCtrl
	msgCtrl  *messageCtrl
}

// New creates a new instance of API type
func New() *API {
	api := new(API)
	api.msgCtrl = newMsgCtrl()
	api.pingCtrl = newPingCtrl()
	return api
}

type ginCtx interface {
	Param(key string) string

	String(code int, format string, values ...interface{})
	JSON(code int, obj interface{})
}

// Run the API
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
