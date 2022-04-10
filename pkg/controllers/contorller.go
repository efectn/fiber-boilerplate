package controllers

import "github.com/efectn/fiber-boilerplate/pkg/services"

type Controller struct {
	Article *ArticleController
}

func NewController(articleService *services.ArticleService) *Controller {
	return &Controller{
		Article: &ArticleController{articleService: articleService},
	}
}
