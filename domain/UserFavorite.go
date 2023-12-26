package domain

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

//counterfeiter:generate . UserFavoriteRepository
type UserFavoriteRepository interface {
	FindAllByUserId(ctx context.Context, param FindAllUserFavoriteParam) (res FindAllArticleResult, err error)
	UpSertByUserId(ctx context.Context, userFavorite model.UserFavorite) (err error)
	DeleteOneByUserId(ctx context.Context, userId string, articleId string) (err error)
}

type FindAllUserFavoriteParam struct {
	WithTag       bool
	Orders        gdb.OrderByParams
	Pagination    gdb.PaginationParam
	UserId        string
	ArticleFields []string
}
