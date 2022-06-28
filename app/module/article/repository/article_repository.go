package repository

import (
	"context"

	"github.com/efectn/fiber-boilerplate/app/module/article/request"
	"github.com/efectn/fiber-boilerplate/internal/bootstrap/database"
	"github.com/efectn/fiber-boilerplate/internal/ent"
	article "github.com/efectn/fiber-boilerplate/internal/ent/article"
)

type ArticleRepository struct {
	DB *database.Database
}

type IArticleRepository interface {
	GetArticles() ([]*ent.Article, error)
	GetArticleByID(id int) (*ent.Article, error)
	CreateArticle(request request.ArticleRequest) (*ent.Article, error)
	UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error)
	DeleteArticle(id int) error
}

func NewArticleRepository(database *database.Database) *ArticleRepository {
	return &ArticleRepository{
		DB: database,
	}
}

func (s *ArticleRepository) GetArticles() ([]*ent.Article, error) {
	return s.DB.Ent.Article.Query().Order(ent.Asc(article.FieldID)).All(context.Background())
}

func (s *ArticleRepository) GetArticleByID(id int) (*ent.Article, error) {
	return s.DB.Ent.Article.Query().Where(article.IDEQ(id)).First(context.Background())
}

func (s *ArticleRepository) CreateArticle(request request.ArticleRequest) (*ent.Article, error) {
	return s.DB.Ent.Article.Create().
		SetTitle(request.Title).
		SetContent(request.Content).
		Save(context.Background())
}

func (s *ArticleRepository) UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error) {
	return s.DB.Ent.Article.UpdateOneID(id).
		SetTitle(request.Title).
		SetContent(request.Content).
		Save(context.Background())
}

func (s *ArticleRepository) DeleteArticle(id int) error {
	return s.DB.Ent.Article.DeleteOneID(id).Exec(context.Background())
}
