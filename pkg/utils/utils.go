package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

type errorResponse struct {
	Name    string
	Tag     string
	Message string
}

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	enabled := true
	if key {
		enabled = false
	}

	return func(c *fiber.Ctx) bool { return enabled }
}

func ValidateStruct(input interface{}) []*errorResponse {
	// Create validator object
	if validate == nil {
		validate = validator.New()
	}

	var errors []*errorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.Name = err.Field()
			element.Tag = err.Tag()
			element.Message = err.Error()
			errors = append(errors, &element)
		}
	}

	return errors
}

func ParseBody(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}

	return nil
}

func ParseAndValidate(c *fiber.Ctx, body interface{}) []*errorResponse {
	v := reflect.ValueOf(body)

	switch v.Kind() {
	case reflect.Ptr:
		ParseBody(c, body)

		return ValidateStruct(v.Elem().Interface())
	case reflect.Struct:
		ParseBody(c, &body)

		return ValidateStruct(v)
	default:
		return nil
	}
}
