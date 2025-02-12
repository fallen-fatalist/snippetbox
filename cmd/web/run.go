package web

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/repository/postgres"
	"github.com/fallen-fatalist/snippetbox/internal/server"
	"github.com/fallen-fatalist/snippetbox/internal/server/httpserver"
	"github.com/fallen-fatalist/snippetbox/internal/service"
	"github.com/fallen-fatalist/snippetbox/internal/service/serviceinstance"

	_ "github.com/lib/pq"
)

// Actual main
func Run() {
	// Initialize logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize config
	cfg := config.MustConfigLoad()
	logger.Info("Config loaded successfully")

	// Load html templages
	cache, err := httpserver.NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("HTML cache loaded successfully")

	// Open database connection
	db, err := postgres.OpenDB(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	logger.Info("Database connected successfully")

	// Initialize the repositories
	var snippetRepositoryInstance repository.SnippetRepository
	snippetRepositoryInstance, err = postgres.NewSnippetRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Repositories initialized successfully")

	// Initialize the services
	snippetServiceInstance, err := serviceinstance.NewSnippetService(snippetRepositoryInstance)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Services initialized successfully")

	// Initialize general service
	var generalService service.Service
	generalService, err = serviceinstance.NewService(snippetServiceInstance)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("General service initialized successfully")

	// Initialize the app
	var app server.Application = httpserver.NewApp(
		logger,
		cfg,
		generalService,
		cache,
	)

	// Log server start
	logger.Info("Application successfully started", slog.Any("address", app.Config().Port()))

	// Launch server
	err = http.ListenAndServe(":"+app.Config().Port(), app.Routes())

	// In case of error server start log it
	logger.Error(err.Error())
	os.Exit(1)
}
