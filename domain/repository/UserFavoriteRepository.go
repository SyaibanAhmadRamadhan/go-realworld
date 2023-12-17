package repository

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

type UserFavoriteRepository interface {
	FindAllArticleByUserID(ctx context.Context, param ParamFindAllArticleByUserID) (res ResultFindAllArticleByUserID, err error)
	UpSertByUserID(ctx context.Context, userFavorite model.UserFavorite) (err error)
	DeleteOneByUserID(ctx context.Context, userID string, articleID string) (err error)
}

type ResultFindAllArticleByUserID struct {
	Articles []ResultFindOneArticle
	Total    int64 `bson:"total"`
}

type ParamFindAllArticleByUserID struct {
	WithTag       bool
	Orders        gdb.OrderByParams
	Pagination    gdb.PaginationParam
	UserID        string
	ArticleFields []string
}
