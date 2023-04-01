package userRepository

import (
	"encoding/json"
	"errors"
	"github.com/lara-store/damo/app/core/domain"
)

type memoryRepository struct {
	users map[string][]byte
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{
		users: make(map[string][]byte),
	}
}

func (repo *memoryRepository) Get(id string) (domain.User, error) {
	if value, ok := repo.users[id]; ok {
		user := domain.User{}
		err := json.Unmarshal(value, &user)
		if err == nil {
			return user, nil
		}
	}

	return domain.User{}, errors.New("user not found")
}

func (repo *memoryRepository) Save(user *domain.User) error {
	value, err := json.Marshal(user)
	if err == nil {
		repo.users[user.ID] = value
		return nil
	}

	return errors.New("user not save")
}
