package http

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
)

type UseCase interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}
