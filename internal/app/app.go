package app

import (
	"github.com/ValGoldun/clean-template/internal/app/di"
	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(
			di.ProvideConfig,

			di.ProvideUserRepository,
			di.ProvideUserUseCase,
			di.ProvideUserController,

			di.ProvideRouter,

			di.ProvideServer,
		),
		fx.Invoke(
			di.ApplyUserRoutes,
			di.InvokeServer,
		),
	).Run()
}
