package controllers

import (
	"strconv"

	"github.com/efectn/fiber-boilerplate/pkg/models"
	"github.com/efectn/fiber-boilerplate/pkg/utils"
	"github.com/efectn/fiber-boilerplate/pkg/utils/errors"
	"github.com/gofiber/fiber/v2"
)

type ArticleRequest struct {
	Title   string `json:"title" form:"title" validate:"required,max=255"`
	Content string `json:"content" form:"content" validate:"required"`
}

type ArticleController struct{}

var articles = make([]*models.Article, 0)

func (ArticleController) Index(c *fiber.Ctx) error {
	if len(articles) == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(articles)
}

func (ArticleController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	for _, article := range articles {
		if article.ID == id {
			return c.Status(fiber.StatusOK).JSON(article)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Article not found!")
}

func (ArticleController) Store(c *fiber.Ctx) error {
	req := new(ArticleRequest)
	if err := utils.ParseAndValidate(c, req); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	id := 1
	if len(articles) > 0 {
		id = articles[len(articles)-1].ID + 1
	}

	articles = append(articles, &models.Article{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	})

	return c.JSON(articles)
}

func (ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(ArticleRequest)
	if err := utils.ParseAndValidate(c, req); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	for _, article := range articles {
		if article.ID == id {
			article.Title = req.Title
			article.Content = req.Content

			return c.Status(fiber.StatusOK).JSON(article)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Article not found!")
}

func (ArticleController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	i := 0
	for _, article := range articles {
		if article.ID == id {
			ret := make([]*models.Article, 0)
			ret = append(ret, articles[:i]...)
			articles = append(ret, articles[i+1:]...)

			return c.Status(fiber.StatusOK).JSON(articles)
		}
		i++
	}

	return c.Status(fiber.StatusNotFound).SendString("Article not found!")
}
