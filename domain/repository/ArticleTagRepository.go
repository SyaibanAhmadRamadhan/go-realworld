package repository

import (
	"context"

	"realworld-go/domain/model"
)

type ArticleTagRepository interface {
	FindAllDetail(ctx context.Context, param ParamFindAllDetailAT, articleColumns ...string) (res ResultFindAllDetailAT, err error)
	FindOneByArticleID(ctx context.Context, articleID string, articleColumns ...string) (res ResultFindOneAT, err error)
	FindTagPopuler(ctx context.Context, limit int64) (res []ResultPopularTagRes, err error)
	ReplaceAll(ctx context.Context, articleTags []model.ArticleTag) (err error)
}

type ParamFindAllDetailAT struct {
	TagIDs     []string
	OrderBy    OrderBy
	Pagination PaginationParam
}

type ResultFindAllDetailAT struct {
	Articles []ResultFindOneAT
	Total    int64 `bson:"total"`
}

type ResultFindOneAT struct {
	Article model.Article `bson:"article"`
	Tags    []model.Tag   `bson:"tags"`
}

type ResultPopularTagRes struct {
	TagID string `bson:"tagID"`
	Count int64  `bson:"count"`
}
