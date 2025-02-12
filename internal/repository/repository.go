package repository

import "github.com/fallen-fatalist/snippetbox/internal/entities"

type SnippetRepository interface {
	Insert(title, content string, expires int) (id int, err error)
	Get(id int) (entities.Snippet, error)
	Latest() ([]entities.Snippet, error)
}
