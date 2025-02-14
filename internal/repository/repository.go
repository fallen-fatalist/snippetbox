package repository

import (
	"errors"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
)

type SnippetRepository interface {
	Insert(title, content string, expires int) (snippetID int, err error)
	Get(snippetID int) (entities.Snippet, error)
	Latest(int) ([]entities.Snippet, error)
}

type UserRepository interface {
	Insert(name, email, password string) (userID int, err error)
	Get(email string) (entities.User, error)
}

var (
	ErrNilDB    = errors.New("nil database provided to repository")
	ErrNoRecord = errors.New("empty record fetched from database")
)
