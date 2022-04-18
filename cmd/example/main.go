package main

import (
	"go.uber.org/fx"

	"github.com/efectn/fiber-boilerplate/pkg/bootstrap"
	"github.com/efectn/fiber-boilerplate/pkg/controller"
	"github.com/efectn/fiber-boilerplate/pkg/database"
	"github.com/efectn/fiber-boilerplate/pkg/middleware"
	"github.com/efectn/fiber-boilerplate/pkg/repository"
	"github.com/efectn/fiber-boilerplate/pkg/router"
	"github.com/efectn/fiber-boilerplate/pkg/service"
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fx.New(
		// Provide patterns
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewLogger),
		fx.Provide(bootstrap.NewFiber),
		fx.Provide(database.NewDatabase),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(controller.NewController),
		fx.Provide(router.NewRouter),
		repository.NewRepository,
		service.NewService,

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.Init(log.Logger)),
	).Run()
}
