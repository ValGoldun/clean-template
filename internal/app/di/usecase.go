package di

import (
	"github.com/ValGoldun/clean-template/internal/repository"
	"github.com/ValGoldun/clean-template/internal/usecase"
	"github.com/ValGoldun/clean-template/internal/usecase/user"
)

func ProvideUserUseCase(repository repository.User) usecase.User {
	return user.New(repository)
}
