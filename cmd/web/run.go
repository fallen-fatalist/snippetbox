package web

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/server"
	"github.com/fallen-fatalist/snippetbox/internal/server/httpserver"
)

// Actual main
func Run() {
	// Initialize config and logger
	cfg := config.MustConfigLoad()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// Initialize the server
	var server server.Server = httpserver.NewApp(logger, cfg)

	// Log server start
	logger.Info("Starting server", slog.Any("address", server.Config().Port))

	// Launch server
	err := http.ListenAndServe(":"+server.Config().Port, server.Routes())

	// In case of error server start log it
	logger.Error(err.Error())
	os.Exit(1)
}
