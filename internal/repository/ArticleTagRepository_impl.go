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

func (a *articleTagRepositoryImpl) collection() *mongo.Collection {
	articleTag := model.ArticleTag{}
	return a.db.Collection(articleTag.TableName())
}

func (a *articleTagRepositoryImpl) UpSert(ctx context.Context, articleTag model.ArticleTag) (err error) {
	res, err := a.FindByArticleID(ctx, articleTag.ArticleID)
	if err != nil {
		if !errors.Is(err, repository.ErrDataNotFound) {
			return
		}
	}

	if res.ArticleID == "" {
		_, err = a.collection().InsertOne(ctx, articleTag)
	} else {
		update := bson.D{
			bson.E{Key: "$set", Value: bson.D{
				bson.E{Key: "tagIDs", Value: articleTag.TagIDs},
			}},
		}
		_, err = a.collection().UpdateOne(ctx, bson.D{{Key: "articleID", Value: res.ArticleID}}, update)
	}
	if err != nil {
		return
	}

	return
}

func (a *articleTagRepositoryImpl) FindByTagID(ctx context.Context, tagID string, paginate repository.PaginationParam) (articleTags []model.ArticleTag, total int64, err error) {
	opts := options.Find().SetLimit(paginate.Limit).SetSkip(paginate.Offset)

	filter := bson.D{}
	if tagID != "" {
		filter = append(filter, bson.E{Key: "tagIDs", Value: tagID})
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
	unwindStage := bson.D{bson.E{Key: "$unwind", Value: "$tagIDs"}}
	groupStage := bson.D{
		bson.E{Key: "$group", Value: bson.D{
			bson.E{Key: "_id", Value: "$tagIDs"},
			bson.E{Key: "count", Value: bson.D{
				bson.E{Key: "$sum", Value: 1},
			}},
		}},
	}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "tagIDs", Value: "$_id"},
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

	cur, err := a.collection().Aggregate(ctx, mongo.Pipeline{
		unwindStage, groupStage, projectStage, sortStage, limitStage,
	})

	err = cur.All(ctx, &popularTagRes)

	return
}
