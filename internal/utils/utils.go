package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Name  string
	Tag   string
	Value string
}

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	enabled := true
	if key {
		enabled = false
	}

	return func(c *fiber.Ctx) bool { return enabled }
}

func ValidateStruct(article interface{}) []*errorResponse {
	var errors []*errorResponse
	validate := validator.New()
	err := validate.Struct(article)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.Name = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
