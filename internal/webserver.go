package internal

import (
	"net/http"
	"time"

	"github.com/efectn/fiber-boilerplate/internal/config"
	"github.com/efectn/fiber-boilerplate/internal/routes"
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
	app    *fiber.App
	store  *session.Store
	config *config.Config
}

func SetupWebServer(config *config.Config) (*WebServer, error) {
	// Setup Webserver
	ws := &WebServer{
		app: fiber.New(fiber.Config{
			ServerHeader: config.Webserver.Header,
			AppName:      config.Webserver.AppName,
		}),
		store: session.New(session.Config{
			Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
		}),
		config: config,
	}

	// Add Extra Middlewares
	ws.app.Use(logger.New(logger.Config{
		Next:       utils.IsEnabled(config.Logger.Enabled),
		TimeFormat: config.Logger.Timeformat,
		TimeZone:   config.Logger.Timezone,
		Format:     config.Logger.Format,
	}))

	ws.app.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(config.Limiter.Enabled),
		Max:        config.Limiter.Max,
		Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
	}))

	ws.app.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(config.Compress.Enabled),
		Level: config.Compress.Level,
	}))

	ws.app.Use(recover.New(recover.Config{
		Next: utils.IsEnabled(config.Recover.Enabled),
	}))

	ws.app.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir("./storage/public"),
		Browse: true,
		MaxAge: 3600,
	}))

	// Test Routes
	ws.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	ws.app.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(string(example))
	})

	ws.app.Get("/monitor", monitor.New())

	// Add specific routes
	routes.RegisterAPIRoutes(ws.app)

	return ws, nil
}

func (ws *WebServer) ListenWebServer() error {
	err := ws.app.Listen(ws.config.Webserver.Port)
	if err != nil {
		return err
	}

	return nil
}
