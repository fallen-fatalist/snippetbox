package repository

import "github.com/fallen-fatalist/snippetbox/internal/entities"

type SnippetRepository interface {
	Insert(title, content string, expires int) (snippetID int64, err error)
	Get(snippetID int64) (entities.Snippet, error)
	Latest() ([]entities.Snippet, error)
}
