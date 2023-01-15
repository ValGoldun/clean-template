package main

import (
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/app"
	"log"
)

func main() {
	//читаем конфиг
	cfg, err := config.NewConfig()
	if err != nil {
		//пишем ошибку если не получилось
		log.Fatal(err)
	}

	//запускаем приложение
	app.Run(cfg)
}
