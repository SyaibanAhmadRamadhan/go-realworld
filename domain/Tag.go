package domain

import (
	"context"

	"realworld-go/domain/dto"
	"realworld-go/domain/model"
)

//counterfeiter:generate . TagRepository
type TagRepository interface {
	FindAllByNames(ctx context.Context, tagNames []string) (tags []model.Tag, err error)
	FindByName(ctx context.Context, name string) (tag model.Tag, err error)
	FindTagPopuler(ctx context.Context, limit int64) (res []FindTagPopulerResult, err error)
	UpSertMany(ctx context.Context, tagName []string) (err error)
	DeleteById(ctx context.Context, tag model.Tag) (err error)
}

type FindTagPopulerResult struct {
	TagId string `bson:"tagId"`
	Count int64  `bson:"count"`
}

// USECASE

type TagUsecase interface {
	FindTagPopuler(ctx context.Context, limit int64) (res []dto.ResponseTag, err error)
}
