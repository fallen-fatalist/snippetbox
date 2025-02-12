package config

import (
	"flag"
)

type Config struct {
	Port      string
	StaticDir string
}

func MustConfigLoad() *Config {
	cfg := Config{}

	flag.StringVar(&cfg.Port, "port", "4000", "Server port")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static/", "Path to the static assets directory")

	flag.Parse()

	return &cfg
}
