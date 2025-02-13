package postgres

import (
	"database/sql"
	"errors"

	"github.com/fallen-fatalist/snippetbox/internal/entities"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
)

type snippetRepository struct {
	db *sql.DB
}

func NewSnippetRepository(db *sql.DB) (*snippetRepository, error) {
	if db == nil {
		return nil, repository.ErrNilDB
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
	query := `
		SELECT snippet_id, title, content, created_at, expires 
		FROM snippets
		WHERE expires > NOW() AND snippet_id = $1`

	row := r.db.QueryRow(query, snippetID)

	var snippet entities.Snippet
	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.CreatedAt, &snippet.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return snippet, repository.ErrNoRecord
		} else {
			return snippet, err
		}
	}
	return snippet, nil
}

func (r *snippetRepository) Latest(count int) ([]entities.Snippet, error) {
	query := `
		SELECT snippet_id, title, content, created_at, expires 
		FROM snippets
		WHERE expires > NOW() 
		ORDER BY snippet_id DESC LIMIT $1
	`

	rows, err := r.db.Query(query, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []entities.Snippet

	for rows.Next() {
		var snippet entities.Snippet

		err = rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.CreatedAt, &snippet.ExpiresAt)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
