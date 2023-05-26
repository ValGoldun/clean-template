package router

import (
	"github.com/ValGoldun/clean-template/internal/controller/user"
	"github.com/ValGoldun/clean-template/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	handler := gin.New()

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	return &Router{Engine: handler}
}

func (router *Router) ApplyUserRoutes(useCase usecase.User) {
	router.Engine.Use(gin.Logger())
	router.Engine.Use(gin.Recovery())

	controller := user.New(useCase)

	group := router.Engine.Group("/api/v1/user")
	group.GET("/list", controller.GetUsers)
}
