package service

import (
	"context"

	"github.com/efectn/fiber-boilerplate/pkg/database"
	"github.com/efectn/fiber-boilerplate/pkg/database/ent"
	article "github.com/efectn/fiber-boilerplate/pkg/database/ent/article"
	"github.com/efectn/fiber-boilerplate/pkg/request"
)

type ArticleService struct {
	DB *database.Database
}

func NewArticleService(database *database.Database) *ArticleService {
	return &ArticleService{
		DB: database,
	}
}

func (s *ArticleService) GetArticles() ([]*ent.Article, error) {
	return s.DB.Ent.Article.Query().Order(ent.Asc(article.FieldID)).All(context.Background())
}

func (s *ArticleService) GetArticleByID(id int) (*ent.Article, error) {
	return s.DB.Ent.Article.Query().Where(article.IDEQ(id)).First(context.Background())
}

func (s *ArticleService) CreateArticle(request request.ArticleRequest) (*ent.Article, error) {
	return s.DB.Ent.Article.Create().
		SetTitle(request.Title).
		SetContent(request.Content).
		Save(context.Background())
}

func (s *ArticleService) UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error) {
	return s.DB.Ent.Article.UpdateOneID(id).
		SetTitle(request.Title).
		SetContent(request.Content).
		Save(context.Background())
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.DB.Ent.Article.DeleteOneID(id).Exec(context.Background())
}
