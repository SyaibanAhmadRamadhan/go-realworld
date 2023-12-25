package domain

import (
	"context"

	"realworld-go/domain/model"
)

//counterfeiter:generate . TagRepository
type TagRepository interface {
	FindAllByNames(ctx context.Context, tagNames []string) (tags []model.Tag, err error)
	FindByID(ctx context.Context, id string) (tag model.Tag, err error)
	FindTagPopuler(ctx context.Context, limit int64) (res []FindTagPopulerResult, err error)
	UpSertMany(ctx context.Context, tagName []string) (err error)
	DeleteByID(ctx context.Context, tag model.Tag) (err error)
}

type FindTagPopulerResult struct {
	TagID string `bson:"tagID"`
	Count int64  `bson:"count"`
}
