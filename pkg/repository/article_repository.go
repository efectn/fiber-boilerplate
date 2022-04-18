package repository

import (
	"context"

	"github.com/efectn/fiber-boilerplate/pkg/database"
	"github.com/efectn/fiber-boilerplate/pkg/database/ent"
	article "github.com/efectn/fiber-boilerplate/pkg/database/ent/article"
	"github.com/efectn/fiber-boilerplate/pkg/request"
)

type ArticleRepository struct {
	DB *database.Database
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
