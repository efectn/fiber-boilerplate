package controllers

import (
	"strconv"

	"github.com/efectn/fiber-boilerplate/internal/models"
	"github.com/efectn/fiber-boilerplate/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type articleRequest struct {
	Title   string `json:"title" form:"title" validate:"required,max=255"`
	Content string `json:"content" form:"content" validate:"required"`
}

var articles = make(map[int]models.Article)

func ListArticles(c *fiber.Ctx) error {
	if len(articles) == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(articles)
}

func ShowArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	if _, ok := articles[id]; len(articles) == 0 || !ok {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(articles[id])
}

func CreateNewArticle(c *fiber.Ctx) error {
	req := new(articleRequest)
	if err := c.BodyParser(req); err != nil {
		panic(err)
	}

	errors := utils.ValidateStruct(*req)
	if errors != nil {
		return c.Status(403).JSON(errors)

	}

	articles[models.IDs] = models.Article{
		Title:   req.Title,
		Content: req.Content,
	}
	models.IDs++

	return c.JSON(articles)
}

func UpdateArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	req := new(articleRequest)
	if err := c.BodyParser(req); err != nil {
		panic(err)
	}

	errors := utils.ValidateStruct(*req)
	if errors != nil {
		return c.Status(403).JSON(errors)

	}

	if _, ok := articles[id]; len(articles) == 0 || !ok {
		return c.SendStatus(404)
	}

	articles[id] = models.Article{
		Title:   req.Title,
		Content: req.Content,
	}

	return c.Status(200).JSON(articles[id])
}

func DestroyArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}

	if _, ok := articles[id]; len(articles) == 0 || !ok {
		return c.SendStatus(404)
	}

	delete(articles, id)
	return c.Status(200).JSON(articles)
}
