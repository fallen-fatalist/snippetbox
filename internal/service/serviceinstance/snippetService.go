package serviceinstance

import (
	"strings"

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

func (s *snippetService) CreateSnippet(title, content string, expires int) (int64, service.Validator) {
	validator := service.Validator{
		FieldErrors: map[string]error{},
	}
	// Expires days validation
	validator.CheckField(validator.MinValue(expires, service.MinimumExpiresDays), "expires", service.ErrNegativeExpiresDay)
	validator.CheckField(validator.MaxValue(expires, service.MaximumExpiresDays), "expires", service.ErrExceedMaximumExpiresDays)

	// Trim the space in the title and content
	title, content = strings.TrimSpace(title), strings.TrimSpace(content)

	// Title validation
	validator.CheckField(validator.NotBlank(title), "title", service.ErrBlankTitle)
	validator.CheckField(validator.MaxChars(title, service.MaximumTitleLength), "title", service.ErrExceedMaximumTitleLength)

	// Content validation
	validator.CheckField(validator.NotBlank(content), "content", service.ErrBlankContent)
	validator.CheckField(validator.MaxChars(content, service.MaximumContentLength), "content", service.ErrExceedMaximumContentLength)

	if !validator.Valid() {
		return 0, validator
	}

	snippetID, err := s.repository.Insert(title, content, expires)
	if err != nil {
		validator.AddFieldError("err", err)
		return 0, validator
	}
	return snippetID, validator
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
