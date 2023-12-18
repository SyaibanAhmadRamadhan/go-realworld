package domain

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

type ArticleRepository interface {
	FindAllPaginate(ctx context.Context, param FindAllPaginateArticleParam, articleColumns ...string) (res FindAllArticleResult, err error)
	FindOneByID(ctx context.Context, param FindOneByIDArticleParam, columns ...string) (res FindOneArticleResult, err error)
	Create(ctx context.Context, article model.Article) (err error)
	UpdateByID(ctx context.Context, article model.Article, columns []string) (err error)
	DeleteByID(ctx context.Context, article model.Article) (err error)
}

type FindAllPaginateArticleParam struct {
	TagIDs         []string
	Orders         gdb.OrderByParams
	Pagination     gdb.PaginationParam
	AggregationOpt FindArticleOpt
}

type FindOneByIDArticleParam struct {
	ArticleID      string
	AggregationOpt FindArticleOpt
}

type FindArticleOpt struct {
	Tag      bool
	Favorite bool
}

type FindOneArticleResult struct {
	Article  model.Article `bson:"article"`
	Favorite int64         `bson:"favorite"`
	Tags     []model.Tag   `bson:"tags"`
}

type FindAllArticleResult struct {
	Articles []FindOneArticleResult
	Total    int64 `bson:"total"`
}
