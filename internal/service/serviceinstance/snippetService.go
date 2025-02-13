package serviceinstance

import (
	"strings"
	"unicode/utf8"

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

func (s *snippetService) CreateSnippet(title, content string, expires int) (int64, map[string]error) {
	errs := map[string]error{}
	// Expires days validation
	if expires < 1 {
		errs["expires"] = service.ErrNegativeExpiresDay
	} else if expires > service.MaximumExpiresDays {
		errs["expires"] = service.ErrExceedMaximumExpiresDays
	}

	// Trim the space in the title and content
	title, content = strings.TrimSpace(title), strings.TrimSpace(content)

	// Title validation
	if title == "" {
		errs["title"] = service.ErrBlankTitle
	} else if utf8.RuneCountInString(title) > service.MaximumTitleLength {
		errs["title"] = service.ErrExceedMaximumTitleLength
	}

	// Content validation
	if utf8.RuneCountInString(content) > service.MaximumContentLength {
		errs["content"] = service.ErrExceedMaximumContentLength
	} else if content == "" {
		errs["content"] = service.ErrBlankContent
	}

	if len(errs) > 0 {
		return 0, errs
	}

	snippetID, err := s.repository.Insert(title, content, expires)
	if err != nil {
		errs["err"] = err
		return 0, errs
	}
	return snippetID, nil
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
