package application

import (
	"github.com/lara-store/damo/cmd/httpServer/providers"
)

func (route Engine) load() {

	route.GET("/users/:id", providers.UserHandler.Get)
}
