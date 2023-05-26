package repository

import (
	"context"
	"github.com/ValGoldun/clean-template/internal/entity"
)

type User interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}
