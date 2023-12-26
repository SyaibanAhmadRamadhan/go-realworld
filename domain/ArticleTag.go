package domain

import (
	"context"

	"realworld-go/domain/model"
)

//counterfeiter:generate . ArticleTagRepository
type ArticleTagRepository interface {
	ReplaceAll(ctx context.Context, articleTags []model.ArticleTag) (err error)
	DeleteByArticleId(ctx context.Context, articleId string) (err error)
}
