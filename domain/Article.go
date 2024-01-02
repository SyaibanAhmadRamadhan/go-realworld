package domain

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/dto"
	"realworld-go/domain/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . ArticleRepository
type ArticleRepository interface {
	FindAllPaginate(ctx context.Context, param FindAllPaginateArticleParam, articleColumns ...string) (res FindAllArticleResult, err error)
	FindOneByOneColumn(ctx context.Context, param FindOneByIdArticleParam, columns ...string) (res FindOneArticleResult, err error)
	Create(ctx context.Context, article model.Article) (err error)
	UpdateById(ctx context.Context, article model.Article, columns []string) (err error)
	DeleteById(ctx context.Context, article model.Article) (err error)
}

type FindAllPaginateArticleParam struct {
	TagIds         []string
	Orders         gdb.OrderByParams
	Pagination     gdb.PaginationParam
	AggregationOpt FindArticleOpt
}

type FindOneByIdArticleParam struct {
	Column         gdb.FindByOneColumnParam
	AggregationOpt FindArticleOpt
}

type FindArticleOpt struct {
	Tag      bool
	Favorite bool
	Author   bool
}

type FindOneArticleResult struct {
	Article  model.Article `bson:"article"`
	Favorite int64         `bson:"favorite"`
	Tags     []model.Tag   `bson:"tags"`
	Author   model.User    `bson:"author"`
}

type FindAllArticleResult struct {
	Articles []FindOneArticleResult
	Total    int64 `bson:"total"`
}

// USECASE

type ArticleUsecase interface {
	Create(ctx context.Context, req dto.RequestCreateArticle) (res dto.ResponseArticle, err error)
	Update(ctx context.Context, req dto.RequestUpdateArticle) (res dto.ResponseArticle, err error)
	Delete(ctx context.Context, articleId string) (err error)
	FindOne(ctx context.Context, articleId string) (res dto.ResponseArticle, err error)
	FindAll(ctx context.Context, req dto.RequestFindAllArticle) (res dto.ResponseArticles, err error)
}
