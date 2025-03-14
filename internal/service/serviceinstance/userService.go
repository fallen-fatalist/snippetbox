package serviceinstance

import (
	"errors"
	"unicode/utf8"

	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/service"
	"golang.org/x/crypto/bcrypt"
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

func (s *userService) CreateUser(name, email, password string) (int, service.Validator) {
	validator := service.Validator{FieldErrors: map[string]error{}}

	// Blank values check
	validator.CheckField(validator.NotBlank(name), "name", service.ErrBlankName)

	// Email match the format
	validator.CheckField(validator.Matches(email, service.EmailRegex), "email", service.ErrInvalidEmailFormat)

	validator.CheckField(validator.MinChars(password, service.MinimumPasswordLength), "password", service.ErrShortPassword)

	if validator.Valid() {
		password, err := s.hashPassword(password)
		if err != nil {
			validator.AddFieldError("err", err)
			return 0, validator
		} else if utf8.RuneCountInString(password) != service.HashedPasswordLength {
			validator.AddFieldError("err", service.ErrWhileHashing)
			return 0, validator
		}

		userID, err := s.repository.Insert(name, email, password)
		if err != nil {
			if errors.Is(err, service.ErrDuplicateEmail) {
				validator.AddFieldError("email", service.ErrDuplicateEmail)
				return 0, validator
			} else {
				validator.AddFieldError("err", err)
				return 0, validator
			}
		} else {
			return userID, service.Validator{}
		}
	}

	return 0, validator
}

func (s *userService) Authenticate(email, password string) (int, service.Validator) {
	validator := service.Validator{
		FieldErrors:    map[string]error{},
		NonFieldErrors: []error{},
	}

	// Email match the format
	validator.CheckField(validator.Matches(email, service.EmailRegex), "email", service.ErrInvalidEmailFormat)

	// Password requirements
	validator.CheckField(validator.MinChars(password, service.MinimumPasswordLength), "password", service.ErrShortPassword)

	if !validator.Valid() {
		return 0, validator
	}

	user, err := s.repository.Get(email)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			validator.AddNonFieldError(err)
		} else {
			validator.AddFieldError("err", err)
		}
		return 0, validator

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			validator.AddNonFieldError(service.ErrInvalidCredentials)
		} else {
			validator.AddFieldError("err", err)
		}
		return 0, validator
	}
	return user.ID, service.Validator{}
}

func (s *userService) Exists(userID int) (bool, error) {
	return false, nil
}

func (s *userService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
