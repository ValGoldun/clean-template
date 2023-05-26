package di

import (
	"github.com/ValGoldun/clean-template/internal/controller"
	"github.com/ValGoldun/clean-template/internal/router"
)

func ProvideRouter() *router.Router {
	return router.NewRouter()
}

func ApplyUserRoutes(router *router.Router, controller controller.User) {
	router.ApplyUserRoutes(controller)
}
