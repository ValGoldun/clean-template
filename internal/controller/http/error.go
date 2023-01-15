package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func (c *controller) internalError(ctx *gin.Context, err error, msg string) {
	//реальную ошибку пишем в лог, в ответе ее быть не должно
	c.logger.Error(err)

	//отвечаем статус кодом - 500 и человекочитаемой ошибкой
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{msg})
}
