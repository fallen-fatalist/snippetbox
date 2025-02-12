package repository

import (
	"errors"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
)

type SnippetRepository interface {
	Insert(title, content string, expires int) (snippetID int64, err error)
	Get(snippetID int64) (entities.Snippet, error)
	Latest(int) ([]entities.Snippet, error)
}

var (
	ErrNilDB    = errors.New("nil database provided to repository")
	ErrNoRecord = errors.New("empty record fetched from database")
)
