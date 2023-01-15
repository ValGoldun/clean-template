package http

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
	"github.com/ValGoldun/clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UseCase interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}

type controller struct {
	useCase UseCase
	logger  logger.Interface
}

func NewControllerV1(handler *gin.RouterGroup, useCase UseCase, logger logger.Interface) {
	controller := &controller{
		useCase: useCase,
		logger:  logger,
	}

	//указываем обработчик для /users
	handler.GET("/users", controller.users)
}

func (c *controller) users(ctx *gin.Context) {
	//забираем юзеров из юзкейса
	users, err := c.useCase.GetUsers(ctx.Request.Context())
	if err != nil {
		//обрабатываем ошибку
		c.internalError(ctx, err, "server problem")
		return
	}

	//если все ок - отвечаем статусом 200 и json-ом с юзерами
	ctx.JSON(http.StatusOK, users)
}
