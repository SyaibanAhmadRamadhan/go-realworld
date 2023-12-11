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

type articleTagRepositoryImpl struct {
	db *mongo.Database
}

func NewArticleTagRepositoryImpl(db *mongo.Database) repository.ArticleTagRepository {
	return &articleTagRepositoryImpl{
		db: db,
	}
}

func (a *articleTagRepositoryImpl) Create(ctx context.Context, articleTag model.ArticleTag) (err error) {
	_, err = a.collection().InsertOne(ctx, articleTag)
	if err != nil {
		return
	}

	return
}

func (a *articleTagRepositoryImpl) collection() *mongo.Collection {
	article := model.ArticleTag{}
	return a.db.Collection(article.TableName())
}

func (a *articleTagRepositoryImpl) FindByTagID(ctx context.Context, tagID string, paginate repository.PaginationParam) (articleTags []model.ArticleTag, total int64, err error) {
	opts := options.Find().SetLimit(paginate.Limit).SetSkip(paginate.Offset)

	filter := bson.D{}
	if tagID != "" {
		filter = append(filter, bson.E{Key: "tagID", Value: tagID})
	}

	total, err = a.collection().CountDocuments(ctx, filter)
	if err != nil {
		return
	}

	cur, err := a.collection().Find(ctx, filter, opts)
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		var articleTag model.ArticleTag
		if err = cur.Decode(&articleTag); err != nil {
			return
		}
		articleTags = append(articleTags, articleTag)
	}

	return

}

func (a *articleTagRepositoryImpl) FindByArticleID(ctx context.Context, articleID string) (articleTag model.ArticleTag, err error) {
	filter := bson.D{bson.E{Key: "articleID", Value: articleID}}
	err = a.collection().FindOne(ctx, filter).Decode(&articleTag)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = repository.ErrDataNotFound
		}

		return
	}

	return
}

func (a *articleTagRepositoryImpl) FindTagPopuler(ctx context.Context, limit int64) (popularTagRes []repository.PopularTagRes, err error) {
	unwindStage := bson.D{bson.E{Key: "$unwind", Value: "tagID"}}
	groupStage := bson.D{
		bson.E{Key: "$group", Value: bson.D{
			bson.E{Key: "tagIDs", Value: "$tagID"},
			bson.E{Key: "count", Value: bson.D{
				bson.E{Key: "$sum", Value: 1},
			}},
		}},
	}
	sortStage := bson.D{
		bson.E{Key: "$sort", Value: bson.D{
			bson.E{Key: "count", Value: -1},
		}},
	}
	limitStage := bson.D{
		bson.E{Key: "$limit", Value: 10},
	}

	cur, err := a.collection().Aggregate(ctx, mongo.Pipeline{
		unwindStage, groupStage, sortStage, limitStage,
	})

	err = cur.All(ctx, &popularTagRes)

	return
}
