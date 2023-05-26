package controller

import "github.com/gin-gonic/gin"

type User interface {
	Users(ctx *gin.Context)
}
