package controllers

import (
	"strconv"

	"github.com/efectn/fiber-boilerplate/pkg/requests"
	"github.com/efectn/fiber-boilerplate/pkg/services"
	"github.com/efectn/fiber-boilerplate/pkg/utils"
	"github.com/efectn/fiber-boilerplate/pkg/utils/errors"
	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	articleService *services.ArticleService
}

func NewArticleController(articleService *services.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: articleService,
	}
}

func (con *ArticleController) Index(c *fiber.Ctx) error {
	articles, err := con.articleService.GetArticles()
	if err != nil {
		return err
	}
	return c.JSON(articles)
}

func (con *ArticleController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	article, err := con.articleService.GetArticleByID(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The article retrieved successfully!",
		"article": article,
	})
}

func (con *ArticleController) Store(c *fiber.Ctx) error {
	req := new(requests.ArticleRequest)
	if err := utils.ParseAndValidate(c, req); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	article, err := con.articleService.CreateArticle(*req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The article created successfully!",
		"article": article,
	})
}

func (con *ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(requests.ArticleRequest)
	if err := utils.ParseAndValidate(c, req); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	article, err := con.articleService.UpdateArticle(id, *req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The article updated successfully!",
		"article": article,
	})
}

func (con *ArticleController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.articleService.DeleteArticle(id); err != nil {
		return errors.NewErrors(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "The article deleted successfully!",
	})
}
