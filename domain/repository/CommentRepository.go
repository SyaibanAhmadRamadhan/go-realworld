package repository

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

type CommentRepository interface {
	FindAllByArticleID(ctx context.Context, param ParamFindAllByArticleID, fields ...string) (comments []model.Comment, err error)
	UpSertByID(ctx context.Context, comment model.Comment, fields ...string) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}

type ParamFindAllByArticleID struct {
	ArticleID string
	OrderBy   gdb.OrderByParams
	LastID    string
	Limit     int64
}
