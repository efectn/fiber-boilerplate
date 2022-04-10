package router

import (
	"github.com/efectn/fiber-boilerplate/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App        fiber.Router
	Controller *controllers.Controller
}

func NewRouter(fiber *fiber.App, controller *controllers.Controller) *Router {
	return &Router{
		App:        fiber,
		Controller: controller,
	}
}

// Register routes
func (r *Router) Register() {
	// Define controllers
	articleController := r.Controller.Article

	// Define routes
	r.App.Route("/articles", func(router fiber.Router) {
		router.Get("/", articleController.Index)
		router.Get("/:id", articleController.Show)
		router.Post("/", articleController.Store)
		router.Patch("/:id", articleController.Update)
		router.Delete("/:id", articleController.Destroy)
	})
}
