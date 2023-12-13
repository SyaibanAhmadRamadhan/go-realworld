package repository

import (
	"context"

	"realworld-go/domain/model"
)

type ArticleTagRepository interface {
	UpSert(ctx context.Context, articleTag model.ArticleTag) (err error)
	FindByTagID(ctx context.Context, tagID string, paginate PaginationParam) (articleTags []model.ArticleTag, total int64, err error)
	FindByArticleID(ctx context.Context, articleID string) (articleTag model.ArticleTag, err error)
	FindTagPopuler(ctx context.Context, limit int64) (popularTagRes []PopularTagRes, err error)
}

type PopularTagRes struct {
	TagIDs string `bson:"tagIDs"`
	Count  int64  `bson:"count"`
}
