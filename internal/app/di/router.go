package di

import (
	"github.com/ValGoldun/clean-template/internal/controller/router"
	"github.com/ValGoldun/clean-template/internal/usecase"
)

func ProvideRouter() *router.Router {
	return router.NewRouter()
}

func ApplyUserRoutes(router *router.Router, useCase usecase.User) {
	router.ApplyUserRoutes(useCase)
}
