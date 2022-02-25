package webserver

import (
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/efectn/fiber-boilerplate/pkg/utils"
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	"github.com/efectn/fiber-boilerplate/pkg/utils/errors"
	"github.com/efectn/fiber-boilerplate/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	futils "github.com/gofiber/fiber/v2/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type WebServer struct {
	Fiber  *fiber.App
	Config *config.Config
	Logger zerolog.Logger
}

func SetupApp(config *config.Config) (*WebServer, error) {
	// Setup Webserver
	ws := &WebServer{
		Fiber: fiber.New(fiber.Config{
			ServerHeader: config.App.Name,
			AppName:      config.App.Name,
			Prefork:      config.App.Prefork,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
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
			},
			DisableStartupMessage: true,
			IdleTimeout:           config.App.IdleTimeout * time.Second,
			EnablePrintRoutes:     config.App.PrintRoutes,
		}),
		Config: config,
		Logger: zerolog.Logger{},
	}

	// Add Extra Middlewares
	ws.Fiber.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(config.Middleware.Limiter.Enable),
		Max:        config.Middleware.Limiter.Max,
		Expiration: config.Middleware.Limiter.ExpSecs * time.Second,
	}))

	ws.Fiber.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(config.Middleware.Compress.Enable),
		Level: config.Middleware.Compress.Level,
	}))

	ws.Fiber.Use(recover.New(recover.Config{
		Next: utils.IsEnabled(config.Middleware.Recover.Enable),
	}))

	ws.Fiber.Use(pprof.New(pprof.Config{
		Next: utils.IsEnabled(config.Middleware.Pprof.Enable),
	}))

	ws.Fiber.Use(filesystem.New(filesystem.Config{
		Next:   utils.IsEnabled(config.Middleware.Filesystem.Enable),
		Root:   http.Dir(config.Middleware.Filesystem.Root),
		Browse: config.Middleware.Filesystem.Browse,
		MaxAge: config.Middleware.Filesystem.MaxAge,
	}))

	// Test Routes
	ws.Fiber.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Pong! 👋")
	})

	ws.Fiber.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Status(200).SendString(string(example))
	})

	ws.Fiber.Get(config.Middleware.Monitor.Path, monitor.New(monitor.Config{
		Next: utils.IsEnabled(config.Middleware.Monitor.Enable),
	}))

	return ws, nil
}

func (ws *WebServer) SetupLogger() error {
	zerolog.TimeFieldFormat = ws.Config.Logger.TimeFormat

	if ws.Config.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(ws.Config.Logger.Level)

	ws.Logger = log.Hook(PreforkHook{})

	return nil
}

func (ws *WebServer) ListenWebServer() error {
	// Custom Startup Messages
	host, port := config.ParseAddr(ws.Config.App.Port)
	if host == "" {
		if ws.Fiber.Config().Network == "tcp6" {
			host = "[::1]"
		} else {
			host = "0.0.0.0"
		}
	}

	// ASCII Art
	ascii, err := os.ReadFile("./storage/ascii_art.txt")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(futils.UnsafeString(ascii), "\n") {
		ws.Logger.Info().Msg(line)
	}

	// Information message
	ws.Logger.Info().Msg(ws.Fiber.Config().AppName + " is running at the moment!")

	// Debug informations
	if !ws.Config.App.Production {
		prefork := "Enabled"
		procs := runtime.GOMAXPROCS(0)
		if !ws.Config.App.Prefork {
			procs = 1
			prefork = "Disabled"
		}

		ws.Logger.Debug().Msgf("Version: %s", "-")
		ws.Logger.Debug().Msgf("Host: %s", host)
		ws.Logger.Debug().Msgf("Port: %s", port)
		ws.Logger.Debug().Msgf("Prefork: %s", prefork)
		ws.Logger.Debug().Msgf("Handlers: %d", ws.Fiber.HandlersCount())
		ws.Logger.Debug().Msgf("Processes: %d", procs)
		ws.Logger.Debug().Msgf("PID: %d", os.Getpid())
	}

	// Listen the app (with TLS Support)
	if ws.Config.App.TLS.Enable {
		ws.Logger.Debug().Msg("TLS support has enabled.")

		if err := ws.Fiber.ListenTLS(ws.Config.App.Port, ws.Config.App.TLS.CertFile, ws.Config.App.TLS.KeyFile); err != nil {
			return err
		}
	}

	if err := ws.Fiber.Listen(ws.Config.App.Port); err != nil {
		return err
	}

	return nil
}

func (ws *WebServer) ShutdownApp() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	ws.Logger.Info().Msg("Shutting down the app...")
	if err := ws.Fiber.Shutdown(); err != nil {
		ws.Logger.Panic().Err(err).Msg("")
	}

	ws.Logger.Info().Msg("Running cleanup tasks...")
	ws.Logger.Info().Msgf("%s was successful shutdown.", ws.Config.App.Name)
	ws.Logger.Info().Msg("\u001b[96msee you again👋\u001b[0m")

	os.Exit(1)
}

// Prefork hook for zerolog
type PreforkHook struct{}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if fiber.IsChild() {
		e.Discard()
	}
}
