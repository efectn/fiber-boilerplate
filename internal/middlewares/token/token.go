// Simple token middleware as example

package token

import "github.com/gofiber/fiber/v2"

type Config struct {
	Next       func(c *fiber.Ctx) bool
	Token      string
	HeaderName string
}

var ConfigDefault = Config{
	Next:       nil,
	Token:      "Default",
	HeaderName: "X-Token-Middleware",
}

func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.Token == "" {
		cfg.Token = ConfigDefault.Token
	}
	if cfg.HeaderName == "" {
		cfg.HeaderName = ConfigDefault.HeaderName
	}

	return cfg
}

// New creates a new middleware handler
func New(config Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		if c.Get(cfg.HeaderName, "") == cfg.Token {
			return c.Next()
		}

		return c.SendStatus(fiber.StatusUnauthorized)
	}
}
