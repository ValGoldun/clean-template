package controller

import "github.com/gin-gonic/gin"

type User interface {
	GetUsers(ctx *gin.Context)
}
