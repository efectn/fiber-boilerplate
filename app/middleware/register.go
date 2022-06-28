package middleware

import (
	"net/http"
	"time"

	"github.com/efectn/fiber-boilerplate/utils"
	"github.com/efectn/fiber-boilerplate/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Middleware struct {
	App *fiber.App
	Cfg *config.Config
}

func NewMiddleware(app *fiber.App, cfg *config.Config) *Middleware {
	return &Middleware{
		App: app,
		Cfg: cfg,
	}
}

func (m *Middleware) Register() {
	// Add Extra Middlewares
	m.App.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(m.Cfg.Middleware.Limiter.Enable),
		Max:        m.Cfg.Middleware.Limiter.Max,
		Expiration: m.Cfg.Middleware.Limiter.ExpSecs * time.Second,
	}))

	m.App.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(m.Cfg.Middleware.Compress.Enable),
		Level: m.Cfg.Middleware.Compress.Level,
	}))

	m.App.Use(recover.New(recover.Config{
		Next: utils.IsEnabled(m.Cfg.Middleware.Recover.Enable),
	}))

	m.App.Use(pprof.New(pprof.Config{
		Next: utils.IsEnabled(m.Cfg.Middleware.Pprof.Enable),
	}))

	m.App.Use(filesystem.New(filesystem.Config{
		Next:   utils.IsEnabled(m.Cfg.Middleware.Filesystem.Enable),
		Root:   http.Dir(m.Cfg.Middleware.Filesystem.Root),
		Browse: m.Cfg.Middleware.Filesystem.Browse,
		MaxAge: m.Cfg.Middleware.Filesystem.MaxAge,
	}))

	m.App.Get(m.Cfg.Middleware.Monitor.Path, monitor.New(monitor.Config{
		Next: utils.IsEnabled(m.Cfg.Middleware.Monitor.Enable),
	}))
}
