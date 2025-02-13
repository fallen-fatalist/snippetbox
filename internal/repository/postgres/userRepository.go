package postgres

import (
	"database/sql"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*userRepository, error) {
	if db == nil {
		return nil, repository.ErrNilDB
	}

	return &userRepository{db}, nil
}

func (r *userRepository) Insert(name, email, password string) (userID int, err error) {
	return 0, nil
}

func (r *userRepository) Get(email string) (entities.User, error) {
	return entities.User{}, nil
}
