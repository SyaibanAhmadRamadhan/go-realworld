package repository

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

type ArticleRepository interface {
	FindAllPaginate(ctx context.Context, param ParamFindAllPaginate, articleColumns ...string) (res ResultFindAllArticle, err error)
	FindOneByID(ctx context.Context, param ParamFindOneByID, columns ...string) (res ResultFindOneArticle, err error)
	Create(ctx context.Context, article model.Article) (err error)
	UpdateByID(ctx context.Context, article model.Article, columns []string) (err error)
	DeleteByID(ctx context.Context, article model.Article) (err error)
}

type ParamFindAllPaginate struct {
	TagIDs         []string
	Orders         gdb.OrderByParams
	Pagination     PaginationParam
	AggregationOpt ParamFindAllPaginateOpt
}

type ParamFindOneByID struct {
	ArticleID      string
	AggregationOpt ParamFindAllPaginateOpt
}

type ParamFindAllPaginateOpt struct {
	Tag      bool
	Favorite bool
}

type ResultFindOneArticle struct {
	Article model.Article `bson:"article"`
	Tags    []model.Tag   `bson:"tags"`
}

type ResultFindAllArticle struct {
	Articles []ResultFindOneArticle
	Total    int64 `bson:"total"`
}
