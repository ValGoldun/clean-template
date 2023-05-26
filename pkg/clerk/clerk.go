package clerk

import (
	"github.com/ValGoldun/clean-template/pkg/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type Clerk struct{}

func New() Clerk {
	return Clerk{}
}

func (c Clerk) Problem(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case errors.Error:
		if e.IsCritical() {
			c.serverProblem(ctx, err)
			return
		}
		c.clientProblem(ctx, err)
		return
	default:
		c.serverProblem(ctx, err)
	}
}

func (c Clerk) serverProblem(ctx *gin.Context, err error) {
	log.Println(err.Error())

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Error{Error: "server problem"})
}

func (c Clerk) clientProblem(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, Error{Error: err.Error()})
}
