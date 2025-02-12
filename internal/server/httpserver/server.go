package httpserver

import (
	"database/sql"
	"log/slog"

	"github.com/fallen-fatalist/snippetbox/internal/config"
)

type application struct {
	logger *slog.Logger
	cfg    *config.Config
	db     *sql.DB
}

func NewApp(logger *slog.Logger, cfg *config.Config, db *sql.DB) *application {
	return &application{
		logger: logger,
		cfg:    cfg,
		db:     db,
	}
}

// Getters
func (app *application) Config() *config.Config {
	return app.cfg
}

func (app *application) Logger() *slog.Logger {
	return app.logger
}

func (app *application) DB() *sql.DB {
	return app.db
}
