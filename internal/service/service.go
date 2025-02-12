package service

import (
	"errors"
	"fmt"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
)

type Service interface {
	SnippetService() SnippetService
}

type SnippetService interface {
	GetSnippetByID(id int64) (entities.Snippet, error)
	// Expires is the number of days, in which snippet will be expired
	CreateSnippet(title, content string, expires int) (int64, error)
	// Returns last n snippets
	LatestSnippets(n int) ([]entities.Snippet, error)
}

// Constants
var (
	MaximumExpiresDays        = 730
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
	ErrExceedMaximumDays               = fmt.Errorf("expires day exceed maximum %d days", MaximumExpiresDays)
	ErrExceedMaximumTitleLength        = fmt.Errorf("title length exceed maximum %d title length", MaximumTitleLength)
	ErrExceedMaximumContentLength      = fmt.Errorf("content length exceed maximum %d content length", MaximumContentLength)
	ErrNegativeNumberLastSnippets      = fmt.Errorf("negative number of last snippets provided")
	ErrExceedMaximumLastSnippetsNumber = fmt.Errorf("last snippets number exceed maximum %d last snippets number", MaximumLastSnippetsNumber)
)

var serviceErrors = []error{
	ErrNegativeID,
	ErrNegativeExpiresDay,
	ErrExceedMaximumDays,
	ErrExceedMaximumTitleLength,
	ErrExceedMaximumContentLength,
	ErrNegativeNumberLastSnippets,
}

func IsServiceError(err error) bool {
	for _, e := range serviceErrors {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
