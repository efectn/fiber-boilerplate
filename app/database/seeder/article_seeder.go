package seeds

import (
	"context"

	"github.com/efectn/fiber-boilerplate/internal/ent"
)

type ArticleSeeder struct{}

var articles = []ent.Article{
	{
		Title:   "Example 1",
		Content: "Lorem Ipsum",
	},
	{
		Title:   "Example 2",
		Content: "Dolor Sit Amet",
	},
}

func (ArticleSeeder) Seed(conn *ent.Client) error {
	bulk := make([]*ent.ArticleCreate, len(articles))
	for i, article := range articles {
		bulk[i] = conn.Article.Create().
			SetTitle(article.Title).
			SetContent(article.Content)
	}
	_, err := conn.Article.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (ArticleSeeder) Count(conn *ent.Client) (int, error) {
	return conn.Article.Query().Count(context.Background())
}
