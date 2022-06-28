package service

import (
	"github.com/efectn/fiber-boilerplate/app/module/article/repository"
	"github.com/efectn/fiber-boilerplate/app/module/article/request"
	"github.com/efectn/fiber-boilerplate/internal/ent"
)

type ArticleService struct {
	Repo *repository.ArticleRepository
}

type IArticleService interface {
	GetArticles() ([]*ent.Article, error)
	GetArticleByID(id int) (*ent.Article, error)
	CreateArticle(request request.ArticleRequest) (*ent.Article, error)
	UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error)
	DeleteArticle(id int) error
}

func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		Repo: repo,
	}
}

func (s *ArticleService) GetArticles() ([]*ent.Article, error) {
	return s.Repo.GetArticles()
}

func (s *ArticleService) GetArticleByID(id int) (*ent.Article, error) {
	return s.Repo.GetArticleByID(id)
}

func (s *ArticleService) CreateArticle(request request.ArticleRequest) (*ent.Article, error) {
	return s.Repo.CreateArticle(request)
}

func (s *ArticleService) UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error) {
	return s.Repo.UpdateArticle(id, request)
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.Repo.DeleteArticle(id)
}
