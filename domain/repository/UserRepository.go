package repository

import (
	"context"

	"realworld-go/domain/model"
)

type UserRepository interface {
	FindByOneColumn(ctx context.Context, param FindByOneColumnParam, columns ...string) (user model.User, err error)
	Create(ctx context.Context, user model.User) (err error)
	UpdateByID(ctx context.Context, user model.User, columns []string) (err error)
}
