package usecase

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}

type UseCase struct {
	repository Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{
		repository: repo,
	}
}

func (u UseCase) GetUsers(ctx context.Context) ([]entity.User, error) {
	//забираем из базы юзеров, тут может быть какая-то дополнительная логика. Например, в методе создании юзера - формирование uuid
	return u.repository.GetUsers(ctx)
}
