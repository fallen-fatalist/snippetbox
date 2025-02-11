package config

import (
	"flag"
)

type Config struct {
	Address   string
	StaticDir string
}

func MustConfigLoad() *Config {
	cfg := Config{}

	cfg.Address = *flag.String("addr", ":4000", "HTTP network address")
	cfg.StaticDir = *flag.String("static-dir", "./ui/static/", "Path to the static assets directory")

	flag.Parse()

	return &cfg
}
