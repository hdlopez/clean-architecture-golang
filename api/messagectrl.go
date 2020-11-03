package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hdlopez/clean-architecture-golang/apierror"
	"github.com/hdlopez/clean-architecture-golang/message"
)

// MessageCtrl handles message entity
type messageCtrl struct {
	srv messageSrv
}

// messageSrv interface is mainly useful for testing purposes
type messageSrv interface {
	Get(ID string) (*message.Message, error)
}

// New creates a new instance of MessageCtrl (by convention)
func newMsgCtrl() *messageCtrl {
	ctrl := new(messageCtrl)
	ctrl.srv = message.NewService()
	return ctrl
}

func (ctrl *messageCtrl) Get(c *gin.Context) {
	// 1 - Unmashall user parameters from gin.Context
	id := c.Param("id")
	// 2 - Check and handle validation errors (no business validation)

	// 3 - Make what you need with your request. In this case, get the message by ID.
	msg, err := ctrl.srv.Get(id)
	if err != nil {
		apiErr := err.(*apierror.APIError)
		// in case of any errors, build the response properly
		c.JSON(apiErr.Code(), apiErr)
	} else {
		// returns my message.Message object
		c.JSON(http.StatusOK, msg)
	}
	return
}
