package config

import (
	"github.com/pelletier/go-toml"
	"os"
)

type Config struct {
	HTTP struct {
		Addr string
	}
	DB struct {
		Host     string
		Name     string
		User     string
		Password string
	}
}

func NewConfig() (Config, error) {
	cfg := Config{}

	data, err := os.ReadFile("config.toml")
	if err != nil {
		return Config{}, err
	}

	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
