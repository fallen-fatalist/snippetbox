package httpserver

import (
	"log/slog"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

type application struct {
	logger  *slog.Logger
	cfg     *config.Config
	service service.Service
}

func NewApp(logger *slog.Logger, cfg *config.Config, service service.Service) *application {
	return &application{
		logger:  logger,
		cfg:     cfg,
		service: service,
	}
}

// Getters
func (app *application) Config() *config.Config {
	return app.cfg
}

func (app *application) Logger() *slog.Logger {
	return app.logger
}

func (app *application) Service() service.Service {
	return app.service
}
