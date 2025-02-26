package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	port      string
	staticDir string
	db        struct {
		dsn string
		// TODO: Add more tweaks to the PostgreSQL database
		// maxOpenConns
		// maxIdleConns
		// maxIdleTime
	}
	tls bool
	// TODO: Cookie lifetime
}

// Getters
func (c *Config) Port() string {
	return c.port
}

func (c *Config) StaticDir() string {
	return c.staticDir
}

func (c *Config) DSN() string {
	return c.db.dsn
}

func (c *Config) TLS() bool {
	return c.tls
}

func MustConfigLoad() *Config {
	cfg := Config{}

	// DSN components
	var (
		dbHost     string = os.Getenv("DB_HOST")
		dbUser     string = os.Getenv("DB_USER")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName     string = os.Getenv("DB_NAME")
		dbPort     string = os.Getenv("DB_PORT")
	)

	// Server vars
	flag.StringVar(&cfg.port, "port", "4000", "Server port")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to the static assets directory")
	flag.BoolVar(&cfg.tls, "tls", false, "Enables TLS, certificate and key must lie in /tls directory corresponding: /tls/cert.pem and /tls/key.pem")

	// Database vars
	cfg.db.dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	//

	flag.Parse()

	return &cfg
}
