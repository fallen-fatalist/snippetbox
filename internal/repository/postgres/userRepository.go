package postgres

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/service"
	"github.com/lib/pq"
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

func (r *userRepository) Insert(name, email, hashedPassword string) (userID int, err error) {
	query := `
		INSERT INTO users (name, email, hashed_password)
		VALUES($1, $2, $3)
		RETURNING user_id
	`

	row := r.db.QueryRow(query, name, email, hashedPassword)
	if row.Err() != nil {
		// TODO: Test for the following two error cases
		if pqErr, ok := row.Err().(*pq.Error); ok {
			if pqErr.Code == UniqueViolationErrorCode && strings.Contains(pqErr.Message, "users_uc_email") {
				return 0, service.ErrDuplicateEmail
			} else if pqErr.Code == CheckViolationErrorCode && strings.Contains(pqErr.Message, "email_format_check") {
				return 0, service.ErrInvalidEmailFormat
			}
		}
		return 0, row.Err()
	}

	err = row.Scan(&userID)
	return userID, err
}

func (r *userRepository) Get(email string) (entities.User, error) {
	query := `
			SELECT user_id, name, email, hashed_password, created_at 
			FROM users
			WHERE email = $1
		`

	var user entities.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, service.ErrInvalidCredentials
		} else {
			return entities.User{}, err
		}

	}
	return user, nil
}
