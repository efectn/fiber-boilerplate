package utils

import "github.com/gofiber/fiber/v2"

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	enabled := true
	if key {
		enabled = false
	}

	return func(c *fiber.Ctx) bool { return enabled }
}
