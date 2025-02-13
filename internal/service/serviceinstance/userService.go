package serviceinstance

import (
	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) (*userService, error) {
	if repository == nil {
		return nil, service.ErrNilRepository
	}

	return &userService{repository: repository}, nil
}

func (s *userService) CreateUser(name, email, password string) error {
	return nil
}

func (s *userService) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (s *userService) Exists(userID int) (bool, error) {
	return false, nil
}
