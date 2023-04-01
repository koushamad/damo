package userService

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lara-store/damo/app/core/domain"
	"github.com/lara-store/damo/app/core/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func New(repository ports.UserRepository) *userService {
	return &userService{
		userRepository: repository,
	}
}

func (srv *userService) Get(id string) (domain.User, error) {
	user, err := srv.userRepository.Get(id)
	if err != nil {
		return domain.User{}, errors.New("user not found from repository")
	}

	return user, nil
}

func (srv *userService) Create(firstName, lastName, email string) (domain.User, error) {

	user := domain.User{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	err := srv.userRepository.Save(&user)
	if err != nil {
		return domain.User{}, errors.New("user not save from repository")
	}

	return user, nil
}
