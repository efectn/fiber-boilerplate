package controllers

import (
	"strconv"

	"github.com/efectn/fiber-boilerplate/pkg/requests"
	"github.com/efectn/fiber-boilerplate/pkg/services"
	"github.com/efectn/fiber-boilerplate/pkg/utils/response"
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

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article list retreived successfully!"},
		Data:     articles,
	})
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

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article retrieved successfully!"},
		Data:     article,
	})
}

func (con *ArticleController) Store(c *fiber.Ctx) error {
	req := new(requests.ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	article, err := con.articleService.CreateArticle(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was created successfully!"},
		Data:     article,
	})
}

func (con *ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(requests.ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	article, err := con.articleService.UpdateArticle(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was updated successfully!"},
		Data:     article,
	})
}

func (con *ArticleController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.articleService.DeleteArticle(id); err != nil {
		return response.NewErrors(fiber.StatusInternalServerError, err.Error())
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was deleted successfully!"},
	})
}
