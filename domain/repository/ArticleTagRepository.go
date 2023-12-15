package repository

import (
	"context"

	"realworld-go/domain/model"
)

type ArticleTagRepository interface {
	ReplaceAll(ctx context.Context, articleTags []model.ArticleTag) (err error)
}
