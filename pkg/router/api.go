package router

import (
	"github.com/efectn/fiber-boilerplate/pkg/controller"
	"github.com/efectn/fiber-boilerplate/storage"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App        fiber.Router
	Controller *controller.Controller
}

func NewRouter(fiber *fiber.App, controller *controller.Controller) *Router {
	return &Router{
		App:        fiber,
		Controller: controller,
	}
}

// Register routes
func (r *Router) Register() {
	// Define controllers
	articleController := r.Controller.Article

	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	r.App.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(string(example))
	})

	// Define routes
	r.App.Route("/articles", func(router fiber.Router) {
		router.Get("/", articleController.Index)
		router.Get("/:id", articleController.Show)
		router.Post("/", articleController.Store)
		router.Patch("/:id", articleController.Update)
		router.Delete("/:id", articleController.Destroy)
	})
}
