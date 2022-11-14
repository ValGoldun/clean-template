package usecase

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}
