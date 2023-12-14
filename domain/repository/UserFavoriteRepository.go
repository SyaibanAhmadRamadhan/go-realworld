package repository

import (
	"context"

	"realworld-go/domain/model"
)

type UserFavoriteRepository interface {
	FindOne(ctx context.Context, param FindByOneColumnParam) (userFav model.UserFavorite, err error)
	Count(ctx context.Context, param FindByOneColumnParam) (total int64, err error)
	InsertOne(ctx context.Context, userFavorite model.UserFavorite) (err error)
	UpdateOne(ctx context.Context, param FindByOneColumnParam, userFavorite model.UserFavorite) (err error)
	DeleteOne(ctx context.Context, param FindByOneColumnParam) (err error)
}
