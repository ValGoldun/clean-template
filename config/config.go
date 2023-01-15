package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

//модель конфига
type Config struct {
	//часть которая отвечает за параметры сервера
	HTTP struct {
		Port string
	}
	//часть которая отвечает за базу данных
	DB struct {
		Host     string
		Name     string
		User     string
		Password string
	}
	//сюда можно добавлять нужные структуры для работы приложения
}

func NewConfig() (*Config, error) {
	//создаем экземпляр Config
	cfg := &Config{}

	//читаем файл config.toml
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		//возвращаем ошибку, если что-то пошло не так, например нет такого файла
		return nil, err
	}

	//читаем конфиг в переменную cfg
	err = toml.Unmarshal(data, cfg)
	if err != nil {
		//возвращаем ошибку, если что-то пошло не так, например неправильная структура
		return nil, err
	}

	//возвращаем cfg
	return cfg, nil
}
