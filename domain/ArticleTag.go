package domain

import (
	"context"

	"realworld-go/domain/model"
)

//counterfeiter:generate . ArticleTagRepository
type ArticleTagRepository interface {
	ReplaceAll(ctx context.Context, articleTags []model.ArticleTag) (err error)
}
