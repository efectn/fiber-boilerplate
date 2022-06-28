package article

import (
	"github.com/efectn/fiber-boilerplate/app/module/article/controller"
	"github.com/efectn/fiber-boilerplate/app/module/article/repository"
	"github.com/efectn/fiber-boilerplate/app/module/article/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ArticleRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// Register bulkly
var NewArticleModule = fx.Options(
	// Register Repository & Service
	fx.Provide(repository.NewArticleRepository),
	fx.Provide(service.NewArticleService),

	// Regiser Controller
	fx.Provide(controller.NewController),

	// Register Router
	fx.Provide(NewArticleRouter),
)

// Router methods
func NewArticleRouter(fiber *fiber.App, controller *controller.Controller) *ArticleRouter {
	return &ArticleRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (r *ArticleRouter) RegisterArticleRoutes() {
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
