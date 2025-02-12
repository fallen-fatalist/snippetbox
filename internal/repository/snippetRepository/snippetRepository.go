package snippetRepository

import (
	"database/sql"
	"errors"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
)

var ErrNilDB = errors.New("nil database provided to repository")

type snippetRepository struct {
	db *sql.DB
}

func NewSnippetRepository(db *sql.DB) (*snippetRepository, error) {
	if db == nil {
		return nil, ErrNilDB
	}
	return &snippetRepository{
		db: db,
	}, nil
}

func (r *snippetRepository) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

func (r *snippetRepository) Get(id int) (entities.Snippet, error) {
	return entities.Snippet{}, nil
}

func (r *snippetRepository) Latest() ([]entities.Snippet, error) {
	return nil, nil
}
