package di

import (
	"context"
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/router"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func ProvideServer(router *router.Router, cfg config.Config) *http.Server {
	return &http.Server{Addr: cfg.HTTP.Addr, Handler: router.Engine}
}

func InvokeServer(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
