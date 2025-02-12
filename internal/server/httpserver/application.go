package httpserver

import (
	"log/slog"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
)

type application struct {
	logger            *slog.Logger
	cfg               *config.Config
	snippetRepository *repository.SnippetRepository
}

func NewApp(logger *slog.Logger, cfg *config.Config, snippetRepository *repository.SnippetRepository) *application {
	return &application{
		logger:            logger,
		cfg:               cfg,
		snippetRepository: snippetRepository,
	}
}

// Getters
func (app *application) Config() *config.Config {
	return app.cfg
}

func (app *application) Logger() *slog.Logger {
	return app.logger
}

func (app *application) SnippetRepository() *repository.SnippetRepository {
	return app.snippetRepository
}
