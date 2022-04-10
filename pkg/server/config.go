package server

import (
	"github.com/efectn/fiber-boilerplate/pkg/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func NewConfig() *config.Config {
	config, err := config.ParseConfig("example")
	if err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}

	return config
}
