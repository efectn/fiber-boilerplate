package main

import (
	"go.uber.org/fx"

	"github.com/efectn/fiber-boilerplate/pkg/middlewares"
	"github.com/efectn/fiber-boilerplate/pkg/router"
	"github.com/efectn/fiber-boilerplate/pkg/server"
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	fx.New(
		fx.Provide(server.NewLogger),
		fx.Provide(server.NewConfig),
		fx.Provide(server.NewFiber),
		fx.Provide(middlewares.NewMiddleware),
		fx.Provide(router.NewRouter),

		fx.Invoke(server.Register),
		fx.WithLogger(fxzerolog.Init(log.Logger)),
	).Run()
}
