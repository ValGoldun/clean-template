package di

import (
	"context"
	"github.com/ValGoldun/clean-template/config"
	"github.com/ValGoldun/clean-template/internal/controller/router"
	"go.uber.org/fx"
	"log"
	"net"
	"net/http"
	"os"
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
			go func() {
				err = server.Serve(ln)
				if err != nil {
					log.Println(err)
					os.Exit(1)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
