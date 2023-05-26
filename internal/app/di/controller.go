package di

import (
	"github.com/ValGoldun/clean-template/internal/controller"
	"github.com/ValGoldun/clean-template/internal/controller/user"
	"github.com/ValGoldun/clean-template/internal/usecase"
	"github.com/ValGoldun/clean-template/pkg/clerk"
)

func ProvideUserController(useCase usecase.User, clerk clerk.Clerk) controller.User {
	return user.New(useCase, clerk)
}
