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

type articleRepositoryImpl struct {
	db *mongo.Database
}

func NewArticleRepositoryImpl(db *mongo.Database) repository.ArticleRepository {
	return &articleRepositoryImpl{
		db: db,
	}
}

func (a *articleRepositoryImpl) collection() *mongo.Collection {
	article := model.Article{}
	return a.db.Collection(article.TableName())
}

func (a *articleRepositoryImpl) FindAllByIDS(ctx context.Context, ids []string, columns ...string) (
	articles []model.Article, err error) {
	opts := options.Find()

	if columns != nil {
		projection := bson.D{}
		for _, column := range columns {
			projection = append(projection, bson.E{Key: column, Value: 1})
		}
		opts.SetProjection(projection)
	}

	filter := bson.D{
		bson.E{
			Key: "_id",
			Value: bson.D{
				bson.E{
					Key: "$in", Value: ids,
				},
			},
		},
	}

	cur, err := a.collection().Find(ctx, filter, opts)
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		var article model.Article
		if err = cur.Decode(&article); err != nil {
			return
		}
		articles = append(articles, article)
	}

	return
}

func (a *articleRepositoryImpl) FindById(ctx context.Context, id string, columns ...string) (art model.Article, err error) {
	projection := bson.D{}
	if len(columns) > 0 {
		for _, column := range columns {
			projection = append(projection, bson.E{Key: column, Value: 1})
		}
	}

	opts := options.FindOne().SetProjection(projection)

	filter := bson.D{{Key: "_id", Value: id}}

	err = a.collection().FindOne(ctx, filter, opts).Decode(&art)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = repository.ErrDataNotFound
		}
	}

	return
}

func (a *articleRepositoryImpl) Create(ctx context.Context, article model.Article) (err error) {
	_, err = a.collection().InsertOne(ctx, article)
	return
}

func (a *articleRepositoryImpl) UpdateByID(ctx context.Context, article model.Article, columns []string) (err error) {
	filter := bson.D{bson.E{Key: "_id", Value: article.ID}}
	set := bson.D{}
	value := article.GetValuesByColums(columns...)

	for i, column := range columns {
		if column == "_id" {
			continue
		}
		set = append(set, bson.E{Key: column, Value: value[i]})
	}

	update := bson.D{{"$set", set}}

	res, err := a.collection().UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		return repository.ErrUpdateDataNotFound
	}

	return
}

func (a *articleRepositoryImpl) DeleteByID(ctx context.Context, article model.Article) (err error) {
	filter := bson.D{bson.E{Key: "_id", Value: article.ID}}

	res, err := a.collection().DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	if res.DeletedCount == 0 {
		return repository.ErrDelDataNotFound
	}

	return
}
