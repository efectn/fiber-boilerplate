package main

import (
	"go.uber.org/fx"

	"github.com/efectn/fiber-boilerplate/pkg/controllers"
	"github.com/efectn/fiber-boilerplate/pkg/database"
	"github.com/efectn/fiber-boilerplate/pkg/middlewares"
	"github.com/efectn/fiber-boilerplate/pkg/router"
	"github.com/efectn/fiber-boilerplate/pkg/server"
	"github.com/efectn/fiber-boilerplate/pkg/services"
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fx.New(
		fx.Provide(server.NewLogger),
		fx.Provide(server.NewConfig),
		fx.Provide(server.NewFiber),
		fx.Provide(database.NewDatabase),
		services.NewService,
		fx.Provide(middlewares.NewMiddleware),
		fx.Provide(controllers.NewController),
		fx.Provide(router.NewRouter),

		fx.Invoke(server.Register),
		fx.WithLogger(fxzerolog.Init(log.Logger)),
	).Run()
}
