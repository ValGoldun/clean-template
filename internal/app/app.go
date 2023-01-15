package app

import (
	"database/sql"
	"fmt"
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/controller/http/router"
	"github.com/ValGoldun/clean-template/internal/usecase"
	"github.com/ValGoldun/clean-template/internal/usecase/repository"
	"github.com/ValGoldun/clean-template/pkg/httpserver"
	"github.com/ValGoldun/clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	//создаем логгер
	logger := logger.New()

	//создаем подключение к бд
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Name,
	))
	if err != nil {
		//пишем ошибку в лог
		logger.Error(err)

		//убиваем процесс с кодом 1 (все коды кроме 0 - неуспешные, эта информация понадобится позже)
		os.Exit(1)
	}

	//закрываем подключение к базе перед закрытием приложения
	defer db.Close()

	//создаем репозиторий
	repository := repository.New(db)

	//создаем юзкейс с репозиторием
	useCase := usecase.New(repository)

	//создаем экземпляр gin.Engine
	handler := gin.New()

	//создаем новый роутер указывая handler - обработчик запросов, usecase и логгер
	router.NewRouter(handler, useCase, logger)

	//создаем новый сервер указывая handler и одну опцию - порт
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	//содаем канал для отлова сигнала завершения процесса от ОС
	interrupt := make(chan os.Signal, 1)

	//указывем в какой канал писать сигнал завершения процесса от ОС
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	//ждем отлова сигнала от ОС о завершении процесса или ошибки от httpServer
	select {
	case s := <-interrupt:
		logger.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(fmt.Errorf("httpServer.Notify: %w", err))
	}

	//выключаем сервер, прекращаем прием запросов, если какие-то запросы еще обрабатываются - они обработаются в течении shutdownTimeout - по умолчанию 3 секунды, если не успеют - соединение будет сброшено
	//нужно для бесшовной выкатки в k8s, когда у нас есть два инстанса приложения и нам нужно их обновить, например с версии 1.0.0 на 1.0.1
	//k8s посылает одному инстансу сигнал завершения процесса, сервер его обрабатывает, выключает сервер, обрабатывает оставшиеся запросы, в то время как все остальные запросы идут в один инстанс версии 1.0.0
	//далее k8s запускает новый инстанс 1.0.1, посылает сигнал завершения во второй инстанс версии 1.0.0, все новые запросы идут только в 1.0.1, старые запросы дообрабатываются вторым инстансом версии 1.0.0
	//после убивания второго инстанса версии 1.0.0 k8s запускает второй инстанс версии 1.0.1
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("httpServer.Shutdown: %w", err))
	}
}
