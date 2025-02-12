package postgres

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

func (r *snippetRepository) Insert(title, content string, expires int) (snippetID int64, err error) {
	query := `
		INSERT INTO snippets 
		(title, content, expires)
		VALUES ($1, $2, NOW() + ($3 * INTERVAL '1 day'))
		RETURNING snippet_id
	`

	args := []interface{}{title, content, expires}
	row := r.db.QueryRow(query, args...)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err = row.Scan(&snippetID)
	if err != nil {
		return 0, err
	}

	return snippetID, nil
}

func (r *snippetRepository) Get(snippetID int64) (entities.Snippet, error) {
	return entities.Snippet{}, nil
}

func (r *snippetRepository) Latest() ([]entities.Snippet, error) {
	return nil, nil
}
