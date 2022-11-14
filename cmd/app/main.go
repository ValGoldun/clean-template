package main

import (
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
