package usecase

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
)

type UseCase struct {
	repository Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{
		repository: repo,
	}
}

func (u UseCase) GetUsers(ctx context.Context) ([]entity.User, error) {
	return u.repository.GetUsers(ctx)
}
