package user

import (
	"github.com/ValGoldun/clean-template/internal/usecase"
	"github.com/ValGoldun/clean-template/pkg/clerk"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	useCase usecase.User
	clerk   clerk.Clerk
}

func New(useCase usecase.User, clerk clerk.Clerk) Controller {
	return Controller{
		useCase: useCase,
		clerk:   clerk,
	}
}

func (c Controller) GetUsers(ctx *gin.Context) {
	users, err := c.useCase.GetUsers(ctx.Request.Context())
	if err != nil {
		c.clerk.Problem(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
