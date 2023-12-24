package domain

import (
	"context"

	"realworld-go/domain/model"
)

//counterfeiter:generate . TagRepository
type TagRepository interface {
	FindAllByIDS(ctx context.Context, ids []string) (tags []model.Tag, err error)
	FindByID(ctx context.Context, id string) (tag model.Tag, err error)
	FindTagPopuler(ctx context.Context, limit int64) (res []FindTagPopulerResult, err error)
	Create(ctx context.Context, tag model.Tag) (err error)
	UpdateByID(ctx context.Context, tag model.Tag, column []string) (err error)
	DeleteByID(ctx context.Context, tag model.Tag) (err error)
}

type FindTagPopulerResult struct {
	TagID string `bson:"tagID"`
	Count int64  `bson:"count"`
}