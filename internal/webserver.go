package internal

import (
	"net/http"
	"time"

	"github.com/efectn/fiber-boilerplate/internal/config"
	"github.com/efectn/fiber-boilerplate/internal/utils"
	"github.com/efectn/fiber-boilerplate/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type WebServer struct {
	App    *fiber.App
	Store  *session.Store
	Config *config.Config
}

func SetupWebServer(config *config.Config) (*WebServer, error) {
	// Setup Webserver
	ws := &WebServer{
		App: fiber.New(fiber.Config{
			ServerHeader: config.Webserver.Header,
			AppName:      config.Webserver.AppName,
		}),
		Store: session.New(session.Config{
			Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
		}),
		Config: config,
	}

	// Add Extra Middlewares
	ws.App.Use(logger.New(logger.Config{
		Next:       utils.IsEnabled(config.Logger.Enabled),
		TimeFormat: config.Logger.Timeformat,
		TimeZone:   config.Logger.Timezone,
		Format:     config.Logger.Format,
	}))

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

func (ws *WebServer) ListenWebServer() error {
	err := ws.App.Listen(ws.Config.Webserver.Port)
	if err != nil {
		return err
	}

	return nil
}
