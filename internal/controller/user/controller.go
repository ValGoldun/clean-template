package user

import (
	"github.com/ValGoldun/clean-template/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
	useCase usecase.User
}

func New(useCase usecase.User) Controller {
	return Controller{
		useCase: useCase,
	}
}

func (c Controller) GetUsers(ctx *gin.Context) {
	users, err := c.useCase.GetUsers(ctx.Request.Context())
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
