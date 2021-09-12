package routes

import (
	"github.com/efectn/fiber-boilerplate/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app fiber.Router) {
	// Register Article Routes
	articles := app.Group("/articles")
	articles.Get("/", controllers.ListArticles)
	articles.Get("/:id", controllers.ShowArticle)
	articles.Post("/", controllers.CreateNewArticle)
	articles.Patch("/:id", controllers.UpdateArticle)
	articles.Delete("/:id", controllers.DestroyArticle)
}
