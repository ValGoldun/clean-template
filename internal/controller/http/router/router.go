package router

import (
	"github.com/ValGoldun/clean-template/internal/controller/http"
	"github.com/ValGoldun/clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, useCase http.UseCase, logger logger.Interface) {
	//логирование запросов
	handler.Use(gin.Logger())

	//отлов паники, при панике в ответ будет отдаваться 500 статус код, не будет разрыва соединения
	handler.Use(gin.Recovery())

	//создаем группу эндпоинтов /api/v1
	http.NewControllerV1(handler.Group("/api/v1"), useCase, logger)
}
