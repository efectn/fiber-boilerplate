package main

import (
	"github.com/rs/zerolog/log"

	"github.com/efectn/fiber-boilerplate/pkg/routes"
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	"github.com/efectn/fiber-boilerplate/pkg/webserver"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Parse config
	config, err := config.ParseConfig("example")
	if err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}

	// Setup webserver
	ws, err := webserver.SetupWebServer(config)
	if err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}

	// Setup Logger
	ws.SetupLogger()

	// Register Routes
	routes.RegisterAPIRoutes(ws.App)

	// Run webserver
	go func() {
		if err := ws.ListenWebServer(); err != nil {
			ws.Logger.Panic().Err(err).Msg("")
		}
	}()

	// Gracefully shutdown
	ws.ShutdownApp()
}
