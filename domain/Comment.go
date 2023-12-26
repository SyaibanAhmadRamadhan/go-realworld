package domain

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/dto"
	"realworld-go/domain/model"
)

//counterfeiter:generate . CommentRepository
type CommentRepository interface {
	FindAllByArticleId(ctx context.Context, param FindAllCommentParam, fields ...string) (comments []model.Comment, err error)
	Create(ctx context.Context, comment model.Comment) (err error)
	UpdateById(ctx context.Context, comment model.Comment, fields ...string) (err error)
	DeleteById(ctx context.Context, comment model.Comment) (err error)
	DeleteByArticleId(ctx context.Context, articleId string) (err error)
}

type FindAllCommentParam struct {
	ArticleId string
	OrderBy   gdb.OrderByParams
	LastId    string
	Limit     int64
}

// USECASE

type CommentUsecase interface {
	FindAll(ctx context.Context, req dto.RequestFindAllComment) (res dto.ResponseComments, err error)
	Create(ctx context.Context, req dto.RequestCreateComment) (res dto.ResponseComment, err error)
	Update(ctx context.Context, req dto.RequestUpdateComment) (res dto.ResponseComment, err error)
	Delete(ctx context.Context, req dto.RequestDeleteComment) (err error)
}
