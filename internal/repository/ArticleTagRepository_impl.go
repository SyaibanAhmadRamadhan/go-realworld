package repository

import (
	"context"
	"fmt"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

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

func (a *articleTagRepositoryImpl) articleColl() *mongo.Collection {
	article := model.Article{}
	return a.db.Collection(article.TableName())
}

func (a *articleTagRepositoryImpl) articleTagColl() *mongo.Collection {
	articleTag := model.ArticleTag{}
	return a.db.Collection(articleTag.TableName())
}

func (a *articleTagRepositoryImpl) tagColl() *mongo.Collection {
	article := model.Tag{}
	return a.db.Collection(article.TableName())
}

func (a *articleTagRepositoryImpl) FindAllDetail(ctx context.Context, param repository.ParamFindAllDetailAT, articleColumns ...string) (res repository.ResultFindAllDetailAT, err error) {
	pipeline := mongo.Pipeline{}

	if len(param.TagIDs) > 0 {
		pipeline = append(pipeline, bson.D{
			{"$match", bson.D{
				{"tagID", bson.D{{"$in", param.TagIDs}}},
			}},
		})
	}

	pipeline = append(pipeline,
		// set group
		bson.D{bson.E{Key: "$group", Value: bson.D{
			bson.E{Key: "_id", Value: "$articleID"},
			bson.E{Key: "tagIDs", Value: bson.D{
				bson.E{Key: "$addToSet", Value: "$tagID"},
			}},
		}}},
		// set project from group
		bson.D{bson.E{Key: "$project", Value: bson.D{
			bson.E{Key: "articleID", Value: "$_id"},
			bson.E{Key: "tagIDs", Value: "$tagIDs"},
			bson.E{Key: "_id", Value: 0},
		}}},
		// set lookup article
		bson.D{bson.E{Key: "$lookup", Value: bson.D{
			bson.E{Key: "from", Value: a.articleColl().Name()},
			bson.E{Key: "localField", Value: "articleID"},
			bson.E{Key: "foreignField", Value: "_id"},
			bson.E{Key: "pipeline", Value: bson.A{
				bson.D{bson.E{Key: "$project", Value: a.lookupProjectionArticle(articleColumns...)}},
			}},
			bson.E{Key: "as", Value: "article"},
		}}},
		// set lookup tag
		bson.D{bson.E{Key: "$lookup", Value: bson.D{
			bson.E{Key: "from", Value: a.articleColl().Name()},
			bson.E{Key: "localField", Value: "tagIDs"},
			bson.E{Key: "foreignField", Value: "_id"},
			bson.E{Key: "as", Value: "tags"},
		}}},
	)

	curTotal, err := a.articleTagColl().Aggregate(ctx, append(pipeline, bson.D{bson.E{Key: "$count", Value: "total"}}))
	if err != nil {
		fmt.Println("error aggregate", err)
		return
	}
	if curTotal.Next(ctx) {
		totalMap := bson.M{}
		if err = curTotal.Decode(&totalMap); err != nil {
			return
		}

		total, ok := totalMap["total"].(int32)
		if !ok {
			return res, repository.ErrInvalidTotalType
		}
		res.Total = int64(total)
	}

	pipeline = append(pipeline,
		bson.D{bson.E{Key: "$sort", Value: bson.D{
			bson.E{Key: "article." + param.OrderBy.Column, Value: gcommon.Ternary(param.OrderBy.IsAscending, 1, -1)},
		}}},
		bson.D{bson.E{Key: "$skip", Value: param.Pagination.Offset}},
		bson.D{bson.E{Key: "$limit", Value: param.Pagination.Limit}},
	)
	cur, err := a.articleTagColl().Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		resData := struct {
			Article []model.Article `bson:"article"`
			Tags    []model.Tag     `bson:"tags"`
		}{}
		if err = cur.Decode(&resData); err != nil {
			return
		}

		var resOne repository.ResultFindOneAT
		resOne.Tags = resData.Tags
		if resData.Article != nil {
			resOne.Article = resData.Article[0]
		}

		res.Articles = append(res.Articles, resOne)
	}

	return
}

func (a *articleTagRepositoryImpl) FindOneByArticleID(ctx context.Context, articleID string, articleColumns ...string) (res repository.ResultFindOneAT, err error) {
	pipeline := mongo.Pipeline{}

	pipeline = append(pipeline,
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "articleID", Value: bson.D{
				{Key: "$eq", Value: articleID},
			}},
		}}},
		// bson.D{{Key: "$limit", Value: 1}},
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$articleID"},
			{Key: "tagIDs", Value: bson.D{
				{Key: "$addToSet", Value: "$tagID"},
			}},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "articleID", Value: "$_id"},
			{Key: "tagIDs", Value: "$tagIDs"},
			bson.E{Key: "_id", Value: 0},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: a.articleColl().Name()},
			{Key: "localField", Value: "articleID"},
			{Key: "foreignField", Value: "_id"},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$project", Value: a.lookupProjectionArticle(articleColumns...)}},
			}},
			{Key: "as", Value: "article"},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: a.tagColl().Name()},
			{Key: "localField", Value: "tagIDs"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "tags"},
		}}},
		bson.D{{Key: "$unwind", Value: "$article"}},
	)

	cur, err := a.articleTagColl().Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	if cur.Next(ctx) {
		resData := struct {
			Article model.Article `bson:"article"`
			Tags    []model.Tag   `bson:"tags"`
		}{}
		if err = cur.Decode(&resData); err != nil {
			return
		}
		res.Article = resData.Article
		res.Tags = resData.Tags
	}

	return
}

func (a *articleTagRepositoryImpl) FindTagPopuler(ctx context.Context, limit int64) (res []repository.ResultPopularTagRes, err error) {
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

	cur, err := a.articleTagColl().Aggregate(ctx, mongo.Pipeline{
		groupStage, projectStage, sortStage, limitStage,
	})

	err = cur.All(ctx, &res)

	return
}

func (a *articleTagRepositoryImpl) ReplaceAll(ctx context.Context, articleTags []model.ArticleTag) (err error) {
	if articleTags == nil {
		return
	}
	_, err = a.articleTagColl().DeleteMany(ctx, bson.D{{Key: "articleID", Value: articleTags[0].ArticleID}})
	if err != nil {
		return
	}

	for _, articleTag := range articleTags {
		_, err = a.articleTagColl().InsertOne(ctx, articleTag)
		if err != nil {
			break
		}
	}

	return
}

func (a *articleTagRepositoryImpl) lookupProjectionArticle(articleColumns ...string) bson.D {
	lookupProjection := bson.D{}
	if articleColumns != nil {
		for _, column := range articleColumns {
			lookupProjection = append(lookupProjection, bson.E{Key: column, Value: 1})
		}
	} else {
		article := model.NewArticle()
		for _, column := range article.AllField() {
			lookupProjection = append(lookupProjection, bson.E{Key: column, Value: 1})
		}
	}
	return lookupProjection
}
