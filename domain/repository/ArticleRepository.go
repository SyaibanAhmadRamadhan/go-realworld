package repository

import (
	"context"

	"realworld-go/domain/model"
)

type ArticleRepository interface {
	FindAllByTag(ctx context.Context, paginate PaginationParam, tag string, columns ...string) (
		articles []model.Article, total int64, err error)
	FindById(ctx context.Context, id int, columns ...string) (model.Article, error)
	Create(ctx context.Context, article model.Article) (err error)
	UpdateByID(ctx context.Context, article model.Article, columns ...string) (err error)
	DeleteByID(ctx context.Context, article model.Article, columns ...string) (err error)
}
