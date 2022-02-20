package routes

import (
	"github.com/efectn/fiber-boilerplate/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Define controllers
var articleController = new(controllers.ArticleController)

// Register routes
func RegisterAPIRoutes(app fiber.Router) {
	app.Route("/articles", func(router fiber.Router) {
		router.Get("/", articleController.Index)
		router.Get("/:id", articleController.Show)
		router.Post("/", articleController.Store)
		router.Patch("/:id", articleController.Update)
		router.Delete("/:id", articleController.Destroy)
	})
}
