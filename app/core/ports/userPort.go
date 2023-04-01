package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/lara-store/damo/app/core/domain"
)

type UserRepository interface {
	Get(string string) (domain.User, error)
	Save(user *domain.User) error
}

type UserService interface {
	Get(string string) (domain.User, error)
	Create(firstName, lastName, email string) (domain.User, error)
}

type UserHandler interface {
	Get(c *gin.Context)
}
