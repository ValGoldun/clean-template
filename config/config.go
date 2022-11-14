package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type Config struct {
	HTTP struct {
		Port string
	}
	DB struct {
		Host     string
		Name     string
		User     string
		Password string
	}
	Log struct {
		Level string
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
