package web

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"

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
func Main() {
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

	// TODO: change for gorilla/sessions
	// TODO: if tls use Secure header enabled
	// Session manager
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	// Initialize the repositories
	var snippetRepositoryInstance repository.SnippetRepository
	snippetRepositoryInstance, err = postgres.NewSnippetRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	var userRepositoryInstance repository.UserRepository
	userRepositoryInstance, err = postgres.NewUserRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Repositories initialized successfully")

	// Initialize the services
	snippetServiceInstance, err := serviceinstance.NewSnippetService(snippetRepositoryInstance)
	if err != nil {
		log.Fatal(err)
	}

	userServiceInstance, err := serviceinstance.NewUserService(userRepositoryInstance)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Services initialized successfully")

	// Initialize general service
	var generalService service.Service
	generalService, err = serviceinstance.NewService(
		snippetServiceInstance,
		userServiceInstance,
	)

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
		sessionManager,
	)

	// Prepare server
	srv := http.Server{
		Addr:     ":" + app.Config().Port(),
		Handler:  app.Routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Log server start
	logger.Info("Application successfully started",
		slog.Any("address", app.Config().Port()),
		slog.Any("tls", app.Config().TLS()),
	)

	// TODO: Change optional TLS server launching
	if cfg.TLS() {
		err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	} else {
		err = srv.ListenAndServe()
	}

	// In case of error server start log it
	logger.Error(err.Error())
	os.Exit(1)
}
