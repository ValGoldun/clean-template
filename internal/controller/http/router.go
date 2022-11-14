package http

import (
	"github.com/ValGoldun/clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, useCase UseCase, logger logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	group := handler.Group("/api/v1")
	{
		newController(group, useCase, logger)
	}
}
