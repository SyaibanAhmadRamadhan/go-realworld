package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

type tagRepositoryImpl struct {
	db *mongo.Database
}

func NewTagRepositoryImpl(db *mongo.Database) domain.TagRepository {
	return &tagRepositoryImpl{
		db: db,
	}
}

func (t *tagRepositoryImpl) FindAllByIDS(ctx context.Context, ids []string) (tags []model.Tag, err error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$in", Value: ids}}},
	}

	cur, err := t.db.Collection(model.TagTableName).Find(ctx, filter)
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

	err = t.db.Collection(model.TagTableName).FindOne(ctx, filter).Decode(&tag)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = domain.ErrDataNotFound
		}
	}

	return
}

func (t *tagRepositoryImpl) FindTagPopuler(ctx context.Context, limit int64) (res []domain.FindTagPopulerResult, err error) {
	groupStage := bson.D{
		bson.E{Key: "$group", Value: bson.D{
			bson.E{Key: "_id", Value: "$tagID"},
			bson.E{Key: "count", Value: bson.D{
				bson.E{Key: "$sum", Value: 1},
			}},
		}},
	}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "tagID", Value: "$_id"},
			{Key: "count", Value: "$count"},
			{Key: "_id", Value: 0},
		}},
	}
	sortStage := bson.D{
		bson.E{Key: "$sort", Value: bson.D{
			bson.E{Key: "count", Value: -1},
		}},
	}
	limitStage := bson.D{
		bson.E{Key: "$limit", Value: limit},
	}

	cur, err := t.db.Collection(model.ArticleTagTableName).Aggregate(ctx, mongo.Pipeline{
		groupStage, projectStage, sortStage, limitStage,
	})

	err = cur.All(ctx, &res)

	return
}

func (t *tagRepositoryImpl) Create(ctx context.Context, tag model.Tag) (err error) {
	_, err = t.db.Collection(model.TagTableName).InsertOne(ctx, tag)
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

	res, err := t.db.Collection(model.TagTableName).UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		err = domain.ErrUpdateDataNotFound
	}

	return
}

func (t *tagRepositoryImpl) DeleteByID(ctx context.Context, tag model.Tag) (err error) {
	filter := bson.D{bson.E{Key: "_id", Value: tag.ID}}

	res, err := t.db.Collection(model.TagTableName).DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	if res.DeletedCount == 0 {
		return domain.ErrDelDataNotFound
	}

	return
}
