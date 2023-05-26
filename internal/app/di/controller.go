package di

import (
	"github.com/ValGoldun/clean-template/internal/controller"
	"github.com/ValGoldun/clean-template/internal/controller/user"
	"github.com/ValGoldun/clean-template/internal/usecase"
)

func ProvideUserController(useCase usecase.User) controller.User {
	return user.New(useCase)
}
