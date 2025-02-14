package httpserver

import (
	"html/template"
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

type application struct {
	logger         *slog.Logger
	cfg            *config.Config
	service        service.Service
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
}

func NewApp(logger *slog.Logger, cfg *config.Config, service service.Service, cache map[string]*template.Template, sessionManager *scs.SessionManager) *application {
	return &application{
		logger:         logger,
		cfg:            cfg,
		service:        service,
		templateCache:  cache,
		sessionManager: sessionManager,
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
