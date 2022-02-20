package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/efectn/fiber-boilerplate/pkg/routes"
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	"github.com/efectn/fiber-boilerplate/pkg/webserver"
	"github.com/gofiber/fiber/v2"
)

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n\u001b[96msee you again👋\u001b[0m")
		os.Exit(1)
	}()
}

func main() {
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
	if err := ws.ListenWebServer(); err != nil {
		ws.Logger.Panic().Err(err).Msg("")
	}
}
