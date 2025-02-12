package server

import (
	"log/slog"
	"net/http"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

type Application interface {
	// Endpoints for handling
	Home(w http.ResponseWriter, r *http.Request)
	SnippetCreate(w http.ResponseWriter, r *http.Request)
	SnippetView(w http.ResponseWriter, r *http.Request)

	// Config
	Config() *config.Config
	// Router
	Routes() *http.ServeMux
	// Logger
	Logger() *slog.Logger
	// Service
	Service() service.Service
}
