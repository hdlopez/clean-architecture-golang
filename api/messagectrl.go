package api

import (
	"net/http"

	"github.com/hdlopez/clean-architecture-golang/apierror"
	"github.com/hdlopez/clean-architecture-golang/message"
)

type MessageController interface {
	Get(c Ctx)
}

// messageCtrl handles message entity
type messageCtrl struct {
	srv message.Service
}

// NewMessageController creates a new instance of messageCtrl
func NewMessageController(srv message.Service) MessageController {
	ctrl := new(messageCtrl)
	ctrl.srv = srv
	return ctrl
}

func (ctrl *messageCtrl) Get(c Ctx) {
	// 1 - Unmashall user parameters from gin.Context
	id := c.Param("id")
	// 2 - Check and handle validation errors (no business validation)

	// 3 - Make what you need with your request. In this case, get the message by ID.
	msg, err := ctrl.srv.Get(id)

	if err != nil {
		apiErr, ok := err.(*apierror.APIError)
		if ok {
			// in case of known API errors, build the response properly
			c.JSON(apiErr.Code(), apiErr)
			return
		}

		c.JSON(http.StatusInternalServerError, apierror.FromError(err))
		return
	}

	// returns my message.Message object
	c.JSON(http.StatusOK, msg)
	return
}
