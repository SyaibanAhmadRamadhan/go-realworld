package repository

import (
	"context"

	"realworld-go/domain/model"
)

type TagRepository interface {
	FindAllByIDS(ctx context.Context, ids []string) (tags []model.Tag, err error)
	FindByID(ctx context.Context, id string) (tag model.Tag, err error)
	Create(ctx context.Context, tag model.Tag) (err error)
	UpdateByID(ctx context.Context, tag model.Tag, column []string) (err error)
	DeleteByID(ctx context.Context, tag model.Tag) (err error)
}
