package services

import (
	"github.com/efectn/fiber-boilerplate/pkg/database"
	"github.com/efectn/fiber-boilerplate/pkg/database/models"
	"github.com/gofiber/fiber/v2"
)

type ArticleService struct {
	Model []*models.Article
}

func (s *ArticleService) GetArticle(c *fiber.Ctx, id int) error {
	for _, article := range database.Articles {
		if article.ID == id {
			return c.Status(fiber.StatusOK).JSON(article)
		}
	}

	return c.Next()
}
