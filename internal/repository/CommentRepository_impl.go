package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

type commentRepositoryImpl struct {
	db *mongo.Database
}

func NewCommentRepositoryImpl(db *mongo.Database) repository.CommentRepository {
	return &commentRepositoryImpl{
		db: db,
	}
}

func (c *commentRepositoryImpl) FindAllByArticleID(ctx context.Context, param repository.ParamFindAllByArticleID, fields ...string) (comments []model.Comment, err error) {
	opts := options.Find().SetLimit(param.Limit).SetSort(bson.D{{Key: "_id", Value: -1}})
	if fields != nil {
		projectStage := bson.D{}
		for _, field := range fields {
			projectStage = append(projectStage, bson.E{Key: field, Value: 1})
		}
		opts.SetProjection(projectStage)
	}

	filter := bson.D{{Key: "articleID", Value: param.ArticleID}}
	if param.LastID != "" {
		filter = append(filter, bson.E{Key: "$lt", Value: param.LastID})
	}

	cur, err := c.db.Collection(model.CommentTableName).Find(ctx, filter, opts)
	if err != nil {
		return
	}
	err = cur.All(ctx, &comments)

	return
}

func (c *commentRepositoryImpl) UpSertByID(ctx context.Context, comment model.Comment, fields ...string) (err error) {
	set := bson.D{}

	if fields != nil {
		values := comment.GetValuesByColums(fields...)
		for i, field := range fields {
			if field == "_id" {
				continue
			}
			set = append(set, bson.E{Key: field, Value: values[i]})
		}
	} else {
		fields = comment.AllField()
		values := comment.GetValuesByColums(fields...)
		for i, field := range fields {
			if field == "_id" {
				continue
			}
			set = append(set, bson.E{Key: field, Value: values[i]})
		}
	}

	updated := bson.D{{Key: "$set", Value: set}}

	opts := options.Update().SetUpsert(true)
	_, err = c.db.Collection(model.CommentTableName).UpdateOne(ctx, nil, updated, opts)

	return
}

func (c *commentRepositoryImpl) DeleteByID(ctx context.Context, id string) (err error) {
	res, err := c.db.Collection(model.CommentTableName).DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if res.DeletedCount == 0 {
		err = repository.ErrDelDataNotFound
	}

	return
}
