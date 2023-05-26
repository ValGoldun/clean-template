package di

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/repository"
	"github.com/ValGoldun/clean-template/internal/repository/user"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func ProvideUserRepository(lc fx.Lifecycle, config config.Config) (repository.User, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Name,
	))
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		return db.Close()
	}})

	return user.New(db), nil
}
