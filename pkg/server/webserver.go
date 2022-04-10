package server

import (
	"context"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/efectn/fiber-boilerplate/pkg/database"
	"github.com/efectn/fiber-boilerplate/pkg/middlewares"
	"github.com/efectn/fiber-boilerplate/pkg/router"
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	"github.com/efectn/fiber-boilerplate/pkg/utils/errors"
	"github.com/efectn/fiber-boilerplate/storage"
	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var errorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var messages interface{}

	if e, ok := err.(*errors.Error); ok {
		code = e.Code
		messages = e.Message
	}

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		messages = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"status":   false,
		"messages": messages,
	})
}

func NewFiber(cfg *config.Config) *fiber.App {
	// Setup Webserver
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          errorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
	})

	// Test Routes
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Pong! 👋")
	})

	app.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Status(200).SendString(string(example))
	})

	return app
}

func Register(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, middlewares *middlewares.Middleware, database *database.Database) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddr(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// ASCII Art
				ascii, err := os.ReadFile("./storage/ascii_art.txt")
				if err != nil {
					log.Debug().Err(err).Msg("An unknown error occured when to print ASCII art!")
				}

				for _, line := range strings.Split(futils.UnsafeString(ascii), "\n") {
					log.Info().Msg(line)
				}

				// Information message
				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug().Msgf("Version: %s", "-")
					log.Debug().Msgf("Host: %s", host)
					log.Debug().Msgf("Port: %s", port)
					log.Debug().Msgf("Prefork: %s", prefork)
					log.Debug().Msgf("Handlers: %d", fiber.HandlersCount())
					log.Debug().Msgf("Processes: %d", procs)
					log.Debug().Msgf("PID: %d", os.Getpid())
				}

				// Listen the app (with TLS Support)
				if cfg.App.TLS.Enable {
					log.Debug().Msg("TLS support was enabled.")

					if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
						log.Error().Err(err).Msg("An unknown error occured when to run server!")
					}
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error().Err(err).Msg("An unknown error occured when to run server!")
					}
				}()

				database.ConnectDatabase()
				log.Info().Msg("Connected the database succesfully!")

				database.MigrateModels()
				database.SeedModels()
				log.Info().Msg("Models were migrated and seeded successfully!")

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("")
				}

				log.Info().Msg("Running cleanup tasks...")
				log.Info().Msg("1- Shutdown the database")
				database.ShutdownDatabase()
				log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
				log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
