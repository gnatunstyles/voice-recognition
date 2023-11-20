package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	ApiKey string `env:"ASSEMBLYAI_API_KEY"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config initialization error: %s", err)
	}
	return cfg, nil
}
