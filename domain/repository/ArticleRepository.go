package repository

import (
	"context"

	"realworld-go/domain/model"
)

type ArticleRepository interface {
	FindAllByIDs(ctx context.Context, ids []string, columns ...string) (articles []model.Article, err error)
	FindById(ctx context.Context, id string, columns ...string) (model.Article, error)
	Create(ctx context.Context, article model.Article) (err error)
	UpdateByID(ctx context.Context, article model.Article, columns []string) (err error)
	DeleteByID(ctx context.Context, article model.Article) (err error)
}
