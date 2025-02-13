package service

import (
	"errors"
	"fmt"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
)

type Service interface {
	SnippetService() SnippetService
	UserService() UserService
}

type SnippetService interface {
	GetSnippetByID(id int64) (entities.Snippet, error)
	// Expires is the number of days, in which snippet will be expired
	// Returns id of created snippet and the name of the field matched its corresponding error
	CreateSnippet(title, content string, expires int) (int64, Validator)
	// Returns last n snippets
	LatestSnippets(n int) ([]entities.Snippet, error)
}

type UserService interface {
	CreateUser(name, email, password string) error
	Authenticate(email, password string) (int, error)
	Exists(userID int) (bool, error)
}

// Constants
var (
	MinimumExpiresDays        = 1
	MaximumExpiresDays        = 365
	MaximumTitleLength        = 100
	MaximumContentLength      = 10000
	MaximumLastSnippetsNumber = 100
)

// Errors
var (
	// Init errors
	ErrNilRepository = errors.New("nil repository provided to service")
	ErrNilService    = errors.New("nil service provided to general service")

	// Snippet erros
	ErrNegativeID                      = errors.New("negative or zero id provided")
	ErrNegativeExpiresDay              = errors.New("negative expires day")
	ErrExceedMaximumExpiresDays        = fmt.Errorf("expires day exceed maximum %d days", MaximumExpiresDays)
	ErrBlankTitle                      = errors.New("blank title provided in snippet")
	ErrExceedMaximumTitleLength        = fmt.Errorf("title length exceed maximum %d title length", MaximumTitleLength)
	ErrBlankContent                    = errors.New("blank content provided in snippet")
	ErrExceedMaximumContentLength      = fmt.Errorf("content length exceed maximum %d content length", MaximumContentLength)
	ErrNegativeNumberLastSnippets      = fmt.Errorf("negative number of last snippets provided")
	ErrExceedMaximumLastSnippetsNumber = fmt.Errorf("last snippets number exceed maximum %d last snippets number", MaximumLastSnippetsNumber)

	// User errors
	ErrNoUser             = errors.New("no matching user found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
)

var serviceErrors = []error{
	ErrNegativeID,
	ErrNegativeExpiresDay,
	ErrExceedMaximumExpiresDays,
	ErrBlankTitle,
	ErrExceedMaximumTitleLength,
	ErrBlankContent,
	ErrExceedMaximumContentLength,
	ErrNegativeNumberLastSnippets,
	ErrExceedMaximumLastSnippetsNumber,
	ErrNoUser,
	ErrInvalidCredentials,
	ErrDuplicateEmail,
}

func IsServiceError(err error) bool {
	for _, e := range serviceErrors {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
