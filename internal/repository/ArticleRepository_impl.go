package repository

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/garray"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

type articleRepositoryImpl struct {
	db *mongo.Database
}

func NewArticleRepositoryImpl(db *mongo.Database) domain.ArticleRepository {
	return &articleRepositoryImpl{
		db: db,
	}
}

func (a *articleRepositoryImpl) FindAllPaginate(ctx context.Context, param domain.FindAllPaginateArticleParam, articleColumns ...string) (
	res domain.FindAllArticleResult, err error) {
	pipeline := mongo.Pipeline{}

	pipeline = append(pipeline,
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: model.ArticleTagTableName},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "articleId"},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$project", Value: bson.D{
					{Key: "tagId", Value: 1},
					{Key: "articleId", Value: 1},
				}}},
			}},
			{Key: "as", Value: "article_tag"},
		}}},
	)

	if len(param.TagIds) > 0 {
		pipeline = append(pipeline, bson.D{
			{"$match", bson.D{
				{"article_tag.tagId", bson.D{{"$in", param.TagIds}}},
			}},
		})
	}

	if param.AggregationOpt.Author {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.UserTableName},
				{Key: "localField", Value: "authorId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "author"},
			}}},
		)
	}

	if param.AggregationOpt.Tag {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.TagTableName},
				{Key: "localField", Value: "article_tag.tagId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "tags"},
			}}},
		)
	}

	if param.AggregationOpt.Favorite {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.UserFavoriteTableName},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "articleId"},
				{Key: "as", Value: "userFavorites"},
			}}},
			bson.D{{Key: "$addFields", Value: bson.D{
				{Key: "userFavoritesCount", Value: bson.D{{Key: "$size", Value: "$userFavorites"}}},
			}}},
		)
	}

	pipeline = append(pipeline,
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "tags", Value: "$tags"},
				{Key: "author", Value: "$author"},
				{Key: "favorite", Value: "$userFavoritesCount"},
				{Key: "article", Value: a.projectionArticle(true, articleColumns...)},
			}},
		},
	)

	curTotal, err := a.db.Collection(model.ArticleTableName).Aggregate(ctx, append(pipeline, bson.D{bson.E{Key: "$count", Value: "total"}}))
	if err != nil {
		return
	}
	if curTotal.Next(ctx) {
		totalMap := bson.M{}
		if err = curTotal.Decode(&totalMap); err != nil {
			return
		}

		total, ok := totalMap["total"].(int32)
		if !ok {
			return res, ErrInvalidTotalType
		}
		res.Total = int64(total)
	}

	orderArticles := param.Orders.FilterDifferent(model.NewArticle().OrderFields())
	if orderArticles != nil {
		sort := bson.D{}
		for _, orderArticle := range orderArticles {
			sort = append(sort,
				bson.E{Key: "article." + orderArticle.Column, Value: gcommon.Ternary(orderArticle.IsAscending, 1, -1)},
			)
		}
		pipeline = append(pipeline,
			bson.D{bson.E{Key: "$sort", Value: sort}},
		)
	}

	pipeline = append(pipeline,
		bson.D{bson.E{Key: "$skip", Value: param.Pagination.Offset}},
		bson.D{bson.E{Key: "$limit", Value: param.Pagination.Limit}},
	)
	cur, err := a.db.Collection(model.ArticleTableName).Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	err = cur.All(ctx, &res.Articles)
	return
}

func (a *articleRepositoryImpl) FindOneById(ctx context.Context, param domain.FindOneByIdArticleParam, articleColumns ...string) (
	res domain.FindOneArticleResult, err error) {
	pipeline := mongo.Pipeline{}

	pipeline = append(pipeline,
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "$eq", Value: param.ArticleId},
			}},
		}}},
	)
	if param.AggregationOpt.Tag || param.AggregationOpt.Favorite {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.ArticleTagTableName},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "articleId"},
				{Key: "as", Value: "article_tag"},
			}}},
		)
	}

	if param.AggregationOpt.Author {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.UserTableName},
				{Key: "localField", Value: "authorId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "author"},
			}}},
		)
	}
	if param.AggregationOpt.Tag {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.TagTableName},
				{Key: "localField", Value: "article_tag.tagId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "tags"},
			}}},
		)
	}
	if param.AggregationOpt.Favorite {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.UserFavoriteTableName},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "articleId"},
				{Key: "as", Value: "userFavorites"},
			}}},
			bson.D{{Key: "$addFields", Value: bson.D{
				{Key: "userFavoritesCount", Value: bson.D{{Key: "$size", Value: "$userFavorites"}}},
			}}},
		)
	}

	project := a.projectionArticle(true, articleColumns...)
	project = garray.AppendUniqueVal(project, bson.E{Key: "_id", Value: "$$ROOT._id"})
	pipeline = append(pipeline,
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "tags", Value: "$tags"},
				{Key: "author", Value: "$author"},
				{Key: "favorite", Value: "$userFavoritesCount"},
				{Key: "_id", Value: 0},
				{Key: "article", Value: project},
			}},
		},
		bson.D{{Key: "$limit", Value: 1}},
	)

	cur, err := a.db.Collection(model.ArticleTableName).Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	if cur.Next(ctx) {
		if err = cur.Decode(&res); err != nil {
			// return
		}

		return
	}

	return res, ErrDataNotFound
}

func (a *articleRepositoryImpl) Create(ctx context.Context, article model.Article) (err error) {
	_, err = a.db.Collection(model.ArticleTableName).InsertOne(ctx, article)
	return
}

func (a *articleRepositoryImpl) UpdateById(ctx context.Context, article model.Article, columns []string) (err error) {
	filter := bson.D{bson.E{Key: "_id", Value: article.Id}}
	set := bson.D{}
	value := article.GetValuesByColums(columns...)

	for i, column := range columns {
		if column == "_id" {
			continue
		}
		set = append(set, bson.E{Key: column, Value: value[i]})
	}

	update := bson.D{{"$set", set}}

	res, err := a.db.Collection(model.ArticleTableName).UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}

	if res.MatchedCount == 0 {
		err = ErrUpdateDataNotFound
	}

	return
}

func (a *articleRepositoryImpl) DeleteById(ctx context.Context, article model.Article) (err error) {
	filter := bson.D{bson.E{Key: "_id", Value: article.Id}}

	res, err := a.db.Collection(model.ArticleTableName).DeleteOne(ctx, filter)
	if err != nil {
		return
	}
	if res.DeletedCount == 0 {
		return ErrDelDataNotFound
	}

	return
}

func (a *articleRepositoryImpl) projectionArticle(root bool, articleColumns ...string) bson.D {
	lookupProjection := bson.D{}
	if articleColumns != nil {
		for _, column := range articleColumns {
			lookupProjection = append(lookupProjection,
				bson.E{Key: column, Value: gcommon.Ternary[any](root, "$$ROOT."+column, 1)})
		}
	} else {
		article := model.NewArticle()
		for _, column := range article.AllField() {
			lookupProjection = append(lookupProjection,
				bson.E{Key: column, Value: gcommon.Ternary[any](root, "$$ROOT."+column, 1)})
		}
	}
	return lookupProjection
}
