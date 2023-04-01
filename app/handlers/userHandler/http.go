package userHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/lara-store/damo/app/core/ports"
)

type HttpHandler struct {
	userService ports.UserService
}

func New(userService ports.UserService) *HttpHandler {
	return &HttpHandler{
		userService: userService,
	}
}

func (hdl *HttpHandler) Get(c *gin.Context) {
	id := c.Param("id")

	user, err := hdl.userService.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
