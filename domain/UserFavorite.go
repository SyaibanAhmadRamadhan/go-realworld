package domain

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

//counterfeiter:generate . UserFavoriteRepository
type UserFavoriteRepository interface {
	FindAllByUserID(ctx context.Context, param FindAllUserFavoriteParam) (res FindAllArticleResult, err error)
	UpSertByUserID(ctx context.Context, userFavorite model.UserFavorite) (err error)
	DeleteOneByUserID(ctx context.Context, userID string, articleID string) (err error)
}

type FindAllUserFavoriteParam struct {
	WithTag       bool
	Orders        gdb.OrderByParams
	Pagination    gdb.PaginationParam
	UserID        string
	ArticleFields []string
}
