package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/efectn/fiber-boilerplate/internal"
	"github.com/efectn/fiber-boilerplate/internal/config"
	"github.com/efectn/fiber-boilerplate/internal/routes"
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
	if err != nil {
		panic(err)
	}

	// Setup webserver
	ws, err := internal.SetupWebServer(config)
	if err != nil {
		panic(err)
	}

	// Register Routes
	routes.RegisterAPIRoutes(ws.App)

	// Run webserver
	ws.ListenWebServer()
}
