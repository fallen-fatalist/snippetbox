package serviceinstance

import (
	"github.com/fallen-fatalist/snippetbox/internal/entities"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

type snippetService struct {
	repository repository.SnippetRepository
}

func NewSnippetService(snippetRepository repository.SnippetRepository) (*snippetService, error) {
	if snippetRepository == nil {
		return nil, service.ErrNilRepository
	}
	return &snippetService{repository: snippetRepository}, nil
}

func (s *snippetService) CreateSnippet(title, content string, expires int) (int64, error) {
	// Validation
	if expires < 1 {
		return 0, service.ErrNegativeExpiresDay
	} else if expires > service.MaximumExpiresDays {
		return 0, service.ErrExceedMaximumDays
	} else if runes := []rune(title); len(runes) > service.MaximumTitleLength {
		return 0, service.ErrExceedMaximumTitleLength
	} else if runes := []rune(content); len(runes) > service.MaximumContentLength {
		return 0, service.ErrExceedMaximumContentLength
	}

	return s.repository.Insert(title, content, expires)
}

func (s *snippetService) GetSnippetByID(snippetID int64) (entities.Snippet, error) {
	if snippetID < 1 {
		return entities.Snippet{}, service.ErrNegativeID
	}

	snippet, err := s.repository.Get(snippetID)
	if err != nil {
		return entities.Snippet{}, err
	}

	return snippet, nil
}

func (s *snippetService) LatestSnippets(n int) ([]entities.Snippet, error) {
	if n < 1 {
		return nil, service.ErrNegativeNumberLastSnippets
	} else if n > service.MaximumLastSnippetsNumber {
		return nil, service.ErrExceedMaximumLastSnippetsNumber
	}

	return s.repository.Latest(n)
}
