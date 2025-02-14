package server

import (
	"log/slog"
	"net/http"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

type Application interface {
	// Snippet Endpoints
	Home(w http.ResponseWriter, r *http.Request)
	SnippetCreate(w http.ResponseWriter, r *http.Request)
	SnippetView(w http.ResponseWriter, r *http.Request)

	// User endpoints
	UserSignup(w http.ResponseWriter, r *http.Request)
	UserLogin(w http.ResponseWriter, r *http.Request)
	UserLogout(w http.ResponseWriter, r *http.Request)

	// Config
	Config() *config.Config
	// Router
	Routes() http.Handler
	// Logger
	Logger() *slog.Logger
	// Service
	Service() service.Service
}
