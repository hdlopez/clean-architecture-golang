package api

import "github.com/gin-gonic/gin"

type pingCtrl struct {
}

func newPingCtrl() *pingCtrl {
	return new(pingCtrl)
}

func (ctrl *pingCtrl) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
