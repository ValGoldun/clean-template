package app

import (
	"database/sql"
	"fmt"
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/controller/http"
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
	logger := logger.New(cfg.Log.Level)

	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Name,
	))
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	repository := repository.New(db)

	useCase := usecase.New(repository)

	handler := gin.New()
	http.NewRouter(handler, useCase, logger)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(fmt.Errorf("httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("httpServer.Shutdown: %w", err))
	}
}
