package router

import (
	"github.com/efectn/fiber-boilerplate/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Define controllers
var articleController = new(controllers.ArticleController)

type Router struct {
	App fiber.Router
}

func NewRouter(fiber *fiber.App) *Router {
	return &Router{
		App: fiber,
	}
}

// Register routes
func (r *Router) Register() {
	r.App.Route("/articles", func(router fiber.Router) {
		router.Get("/", articleController.Index)
		router.Get("/:id", articleController.Show)
		router.Post("/", articleController.Store)
		router.Patch("/:id", articleController.Update)
		router.Delete("/:id", articleController.Destroy)
	})
}
