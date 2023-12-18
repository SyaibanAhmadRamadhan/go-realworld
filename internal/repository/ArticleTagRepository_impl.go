package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

type articleTagRepositoryImpl struct {
	db *mongo.Database
}

func NewArticleTagRepositoryImpl(db *mongo.Database) domain.ArticleTagRepository {
	return &articleTagRepositoryImpl{
		db: db,
	}
}

func (a *articleTagRepositoryImpl) ReplaceAll(ctx context.Context, articleTags []model.ArticleTag) (err error) {
	if articleTags == nil {
		return
	}
	_, err = a.db.Collection(model.ArticleTagTableName).DeleteMany(ctx, bson.D{{Key: "articleID", Value: articleTags[0].ArticleID}})
	if err != nil {
		return
	}

	for _, articleTag := range articleTags {
		_, err = a.db.Collection(model.ArticleTagTableName).InsertOne(ctx, articleTag)
		if err != nil {
			break
		}
	}

	return
}
