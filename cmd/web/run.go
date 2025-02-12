package web

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/fallen-fatalist/snippetbox/internal/config"
	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/repository/snippetRepository"
	"github.com/fallen-fatalist/snippetbox/internal/server"
	"github.com/fallen-fatalist/snippetbox/internal/server/httpserver"

	_ "github.com/lib/pq"
)

// Actual main
func Run() {
	// Initialize config
	cfg := config.MustConfigLoad()

	// Initialize logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Open database connection
	db, err := openDB(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	logger.Info("Database successfully connected")

	// Initialize the repositories
	var snippetRepositoryInstance repository.SnippetRepository
	snippetRepositoryInstance, err = snippetRepository.NewSnippetRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the app
	var app server.Application = httpserver.NewApp(
		logger,
		cfg,
		&snippetRepositoryInstance,
	)

	// Log server start
	logger.Info("Server successfully started", slog.Any("address", app.Config().Port()))

	// Launch server
	err = http.ListenAndServe(":"+app.Config().Port(), app.Routes())

	// In case of error server start log it
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// error.
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return db, err
	}

	// Return the sql.DB connection pool.
	return db, nil
}
