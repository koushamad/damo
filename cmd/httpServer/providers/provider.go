package providers

import (
	"github.com/lara-store/damo/app/core/ports"
	"github.com/lara-store/damo/app/core/services/userService"
	"github.com/lara-store/damo/app/handlers/userHandler"
	"github.com/lara-store/damo/app/repositories/userRepository"
)

var UserRepository ports.UserRepository

var UserService ports.UserService

var UserHandler ports.UserHandler

func init() {
	UserRepository = userRepository.NewMemoryRepository()
	UserService = userService.New(UserRepository)
	UserHandler = userHandler.New(UserService)
}
