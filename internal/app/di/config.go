package di

import (
	"github.com/ValGoldun/clean-template/config"
)

func ProvideConfig() (config.Config, error) {
	return config.NewConfig()
}
