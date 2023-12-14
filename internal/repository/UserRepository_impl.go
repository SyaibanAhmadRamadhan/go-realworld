package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

type userRepositoryImpl struct {
	db *mongo.Database
}

func NewUserRepositoryImpl(db *mongo.Database) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) collection() *mongo.Collection {
	user := model.User{}
	return u.db.Collection(user.TableName())
}
func (u *userRepositoryImpl) FindByOneColumn(ctx context.Context, param repository.FindByOneColumnParam, columns ...string) (user model.User, err error) {
	filter := bson.D{{Key: param.Column, Value: param.Value}}
	projection := bson.D{}
	if columns != nil {
		for _, column := range columns {
			projection = append(projection, bson.E{Key: column, Value: 1})
		}
	}

	opts := options.FindOne().SetProjection(projection)

	err = u.collection().FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = repository.ErrDataNotFound
		}
	}

	return
}

func (u *userRepositoryImpl) Create(ctx context.Context, user model.User) (err error) {
	_, err = u.collection().InsertOne(ctx, user)
	return
}

func (u *userRepositoryImpl) UpdateByID(ctx context.Context, user model.User, columns []string) (err error) {
	set := bson.D{}
	values := user.GetValuesByColums(columns...)

	for i, column := range columns {
		if column == "_id" {
			continue
		}
		set = append(set, bson.E{Key: column, Value: values[i]})
	}

	update := bson.D{{Key: "$set", Value: set}}
	res, err := u.collection().UpdateByID(ctx, user.ID, update)
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		err = repository.ErrUpdateDataNotFound
	}

	return
}