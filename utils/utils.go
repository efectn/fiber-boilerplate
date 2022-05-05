package utils

import (
	"github.com/gofiber/fiber/v2"
)

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c *fiber.Ctx) bool { return true }
}
