package webserver

import (
	"net/http"
	"os"
	"runtime"
	"strings"
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
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	futils "github.com/gofiber/fiber/v2/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type WebServer struct {
	App    *fiber.App
	Store  *session.Store
	Config *config.Config
	Logger zerolog.Logger
}

func SetupWebServer(config *config.Config) (*WebServer, error) {
	// Setup Webserver
	ws := &WebServer{
		App: fiber.New(fiber.Config{
			ServerHeader: config.Webserver.Header,
			AppName:      config.Webserver.AppName,
			Prefork:      config.Webserver.Prefork,
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
		}),
		Store: session.New(session.Config{
			Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
		}),
		Config: config,
		Logger: zerolog.Logger{},
	}

	// Add Extra Middlewares
	ws.App.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(config.Limiter.Enabled),
		Max:        config.Limiter.Max,
		Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
	}))

	ws.App.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(config.Compress.Enabled),
		Level: config.Compress.Level,
	}))

	ws.App.Use(recover.New(recover.Config{
		Next: utils.IsEnabled(config.Recover.Enabled),
	}))

	ws.App.Use(filesystem.New(filesystem.Config{
		Next:   utils.IsEnabled(config.Filesystem.Enabled),
		Root:   http.Dir(config.Filesystem.Root),
		Browse: config.Filesystem.Browse,
		MaxAge: config.Filesystem.MaxAge,
	}))

	// Test Routes
	ws.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Pong! 👋")
	})

	ws.App.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Status(200).SendString(string(example))
	})

	ws.App.Get("/monitor", monitor.New(monitor.Config{
		Next: utils.IsEnabled(config.Monitor.Enabled),
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
	host, port := config.ParseAddr(ws.Config.Webserver.Port)
	if host == "" {
		if ws.App.Config().Network == "tcp6" {
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
	ws.Logger.Info().Msg(ws.App.Config().AppName + " is running at the moment!")

	// Debug informations
	if !ws.Config.Webserver.Production {
		prefork := "Enabled"
		procs := runtime.GOMAXPROCS(0)
		if !ws.Config.Webserver.Prefork {
			procs = 1
			prefork = "Disabled"
		}

		ws.Logger.Debug().Msgf("Version: %s", "-")
		ws.Logger.Debug().Msgf("Host: %s", host)
		ws.Logger.Debug().Msgf("Port: %s", port)
		ws.Logger.Debug().Msgf("Prefork: %s", prefork)
		ws.Logger.Debug().Msgf("Handlers: %d", ws.App.HandlersCount())
		ws.Logger.Debug().Msgf("Processes: %d", procs)
		ws.Logger.Debug().Msgf("PID: %d", os.Getpid())
	}

	if err = ws.App.Listen(ws.Config.Webserver.Port); err != nil {
		return err
	}

	return nil
}

// Prefork hook for zerolog
type PreforkHook struct{}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if fiber.IsChild() {
		e.Discard()
	}
}
