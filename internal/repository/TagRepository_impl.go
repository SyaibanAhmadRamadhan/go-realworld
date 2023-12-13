package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

type tagRepositoryImpl struct {
	db *mongo.Database
}

func NewTagRepositoryImpl(db *mongo.Database) repository.TagRepository {
	return &tagRepositoryImpl{
		db: db,
	}
}

func (t *tagRepositoryImpl) collection() *mongo.Collection {
	article := model.Tag{}
	return t.db.Collection(article.TableName())
}

func (t *tagRepositoryImpl) FindAllByIDS(ctx context.Context, ids []string) (tags []model.Tag, err error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$in", Value: ids}}},
	}

	cur, err := t.collection().Find(ctx, filter)
	for cur.Next(ctx) {
		var tag model.Tag
		if err = cur.Decode(&tag); err != nil {
			return
		}

		tags = append(tags, tag)
	}

	return
}

func (t *tagRepositoryImpl) FindByID(ctx context.Context, id string) (tag model.Tag, err error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: id},
	}

	err = t.collection().FindOne(ctx, filter).Decode(&tag)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = repository.ErrDataNotFound
		}
	}

	return
}

func (t *tagRepositoryImpl) Create(ctx context.Context, tag model.Tag) (err error) {
	_, err = t.collection().InsertOne(ctx, tag)
	return
}

func (t *tagRepositoryImpl) UpdateByID(ctx context.Context, tag model.Tag, columns []string) (err error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: tag.ID},
	}
	set := bson.D{}
	values := tag.GetValuesByColums(columns...)

	for i, column := range columns {
		if column == "_id" {
			continue
		}
		set = append(set, bson.E{Key: column, Value: values[i]})
	}

	update := bson.D{{"$set", set}}

	res, err := t.collection().UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		err = repository.ErrUpdateDataNotFound
	}

	return
}

func (t *tagRepositoryImpl) DeleteByID(ctx context.Context, tag model.Tag) (err error) {
	filter := bson.D{bson.E{Key: "_id", Value: tag.ID}}

	res, err := t.collection().DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	if res.DeletedCount == 0 {
		return repository.ErrDelDataNotFound
	}

	return
}
