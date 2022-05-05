package bootstrap

import (
	"os"

	"github.com/efectn/fiber-boilerplate/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(cfg *config.Config) zerolog.Logger {
	zerolog.TimeFieldFormat = cfg.Logger.TimeFormat

	if cfg.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(cfg.Logger.Level)

	//Commented because of breaking logging in request when to use prefork.
	//log.Logger = log.Hook(PreforkHook{})
	return log.Hook(PreforkHook{})
}

// Prefork hook for zerolog
type PreforkHook struct{}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if fiber.IsChild() {
		e.Discard()
	}
}
