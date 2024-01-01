package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

type commentRepositoryImpl struct {
	db *mongo.Database
}

func NewCommentRepositoryImpl(db *mongo.Database) domain.CommentRepository {
	return &commentRepositoryImpl{
		db: db,
	}
}

func (c *commentRepositoryImpl) FindAllByArticleId(ctx context.Context, param domain.FindAllCommentParam, fields ...string) (comments []model.Comment, err error) {
	opts := options.Find().SetLimit(param.Limit).SetSort(bson.D{{Key: "_id", Value: -1}})
	if fields != nil {
		projectStage := bson.D{}
		for _, field := range fields {
			projectStage = append(projectStage, bson.E{Key: field, Value: 1})
		}
		opts.SetProjection(projectStage)
	}

	filter := bson.D{{Key: "articleId", Value: param.ArticleId}}
	if param.LastId != "" {
		filter = append(filter, bson.E{Key: "$lt", Value: param.LastId})
	}

	cur, err := c.db.Collection(model.CommentTableName).Find(ctx, filter, opts)
	if err != nil {
		return
	}
	err = cur.All(ctx, &comments)

	return
}

func (c *commentRepositoryImpl) Create(ctx context.Context, comment model.Comment) (err error) {
	_, err = c.db.Collection(model.CommentTableName).InsertOne(ctx, comment)

	return
}
func (c *commentRepositoryImpl) UpdateById(ctx context.Context, comment model.Comment, fields ...string) (err error) {
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

	filter := bson.D{
		bson.E{Key: "_id", Value: comment.Id},
		bson.E{Key: "authorId", Value: comment.AuthorId},
		bson.E{Key: "articleId", Value: comment.ArticleId},
	}
	updated := bson.D{{Key: "$set", Value: set}}

	res, err := c.db.Collection(model.CommentTableName).UpdateOne(ctx, filter, updated)

	if res.ModifiedCount == 0 {
		err = ErrUpdateDataNotFound
	}

	return
}

func (c *commentRepositoryImpl) DeleteById(ctx context.Context, comment model.Comment) (err error) {
	filter := bson.D{
		bson.E{Key: "_id", Value: comment.Id},
		bson.E{Key: "authorId", Value: comment.AuthorId},
		bson.E{Key: "articleId", Value: comment.ArticleId},
	}

	res, err := c.db.Collection(model.CommentTableName).DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		err = ErrDelDataNotFound
	}

	return
}

func (c *commentRepositoryImpl) DeleteByArticleId(ctx context.Context, articleId string) (err error) {
	_, err = c.db.Collection(model.CommentTableName).DeleteMany(ctx, bson.D{{Key: "articleId", Value: articleId}})

	return
}
