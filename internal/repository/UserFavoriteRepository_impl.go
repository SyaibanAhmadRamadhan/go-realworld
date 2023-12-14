package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

type userFavoriteRepositoryImpl struct {
	db *mongo.Database
}

func NewUserFavoriteRepositoryImpl(db *mongo.Database) repository.UserFavoriteRepository {
	return &userFavoriteRepositoryImpl{db: db}
}
func (u *userFavoriteRepositoryImpl) collection() *mongo.Collection {
	userFavorite := model.UserFavorite{}
	return u.db.Collection(userFavorite.TableName())
}

func (u *userFavoriteRepositoryImpl) FindOne(ctx context.Context, param repository.FindByOneColumnParam) (userFav model.UserFavorite, err error) {
	// TODO implement me
	panic("implement me")
}

func (u *userFavoriteRepositoryImpl) Count(ctx context.Context, param repository.FindByOneColumnParam) (total int64, err error) {
	// TODO implement me
	panic("implement me")
}

func (u *userFavoriteRepositoryImpl) InsertOne(ctx context.Context, userFavorite model.UserFavorite) (err error) {
	// TODO implement me
	panic("implement me")
}

func (u *userFavoriteRepositoryImpl) UpdateOne(ctx context.Context, param repository.FindByOneColumnParam, userFavorite model.UserFavorite) (err error) {
	// TODO implement me
	panic("implement me")
}

func (u *userFavoriteRepositoryImpl) DeleteOne(ctx context.Context, param repository.FindByOneColumnParam) (err error) {
	// TODO implement me
	panic("implement me")
}
