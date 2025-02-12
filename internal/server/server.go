package server

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/fallen-fatalist/snippetbox/internal/config"
)

type Server interface {
	// Endpoints
	Home(w http.ResponseWriter, r *http.Request)
	SnippetCreate(w http.ResponseWriter, r *http.Request)
	SnippetView(w http.ResponseWriter, r *http.Request)

	// Config
	Config() *config.Config
	// Router
	Routes() *http.ServeMux
	// Logger
	Logger() *slog.Logger
	// Database
	DB() *sql.DB
}
