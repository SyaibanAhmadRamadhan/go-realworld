package domain

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"

	"realworld-go/domain/model"
)

//counterfeiter:generate . UserRepository
type UserRepository interface {
	FindByOneColumn(ctx context.Context, param gdb.FindByOneColumnParam, columns ...string) (user model.User, err error)
	Create(ctx context.Context, user model.User) (err error)
	UpdateById(ctx context.Context, user model.User, columns []string) (err error)
}
