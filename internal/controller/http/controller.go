package http

import (
	"github.com/ValGoldun/clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	useCase UseCase
	logger  logger.Interface
}

func newController(handler *gin.RouterGroup, useCase UseCase, logger logger.Interface) {
	controller := &controller{
		useCase: useCase,
		logger:  logger,
	}

	handler.GET("/users", controller.users)
}

func (c *controller) users(ctx *gin.Context) {
	users, err := c.useCase.GetUsers(ctx.Request.Context())
	if err != nil {
		c.logger.Error(err)
		errorResponse(ctx, http.StatusInternalServerError, "database problems")
		return
	}

	ctx.JSON(http.StatusOK, users)
}
