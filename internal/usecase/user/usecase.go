package user

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
	"github.com/ValGoldun/clean-template/internal/repository"
)

type UseCase struct {
	repository repository.User
}

func New(repo repository.User) *UseCase {
	return &UseCase{
		repository: repo,
	}
}

func (u UseCase) GetUsers(ctx context.Context) ([]entity.User, error) {
	return u.repository.GetUsers(ctx)
}
