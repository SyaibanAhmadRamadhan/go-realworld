package repository

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
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
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "$text", Value: bson.D{{Key: "$search", Value: param.Search}}},
		}}},
	}

	if len(param.TagIds) > 0 {
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.ArticleTagTableName},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "articleId"},
				{Key: "as", Value: "article_tag"},
			}}},
			bson.D{{Key: "$match", Value: bson.D{
				{Key: "article_tag.tagId", Value: bson.D{{Key: "$in", Value: param.TagIds}}},
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
			bson.D{{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$author"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
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

func (a *articleRepositoryImpl) FindOneByOneColumn(ctx context.Context, param domain.FindOneByIdArticleParam, articleColumns ...string) (
	res domain.FindOneArticleResult, err error) {
	pipeline := mongo.Pipeline{}

	// set pipeline match column equal
	pipeline = append(pipeline,
		bson.D{{Key: "$match", Value: bson.D{
			{Key: param.Column.Column, Value: bson.D{
				{Key: "$eq", Value: param.Column.Value},
			}},
		}}},
	)

	// if with aggregation tag true
	if param.AggregationOpt.Tag {
		pipeline = append(pipeline,
			// set lookup article tag collection and unwind result
			bson.D{{"$lookup", bson.D{
				{"from", model.ArticleTagTableName},
				{"localField", "_id"},
				{"foreignField", "articleId"},
				{"as", "article_tags"},
			}}}, bson.D{{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$article_tags"},
			}}},

			// set lookup tag collection and unwind result
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.TagTableName},
				{Key: "let", Value: bson.D{
					{Key: "tagId", Value: "$article_tags.tagId"},
				}},
				{Key: "pipeline", Value: mongo.Pipeline{
					bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{
						{Key: "$eq", Value: bson.A{
							bson.D{{Key: "$toString", Value: "$_id"}},
							"$$tagId",
						}},
					}}}}},
				}},
				{Key: "as", Value: "tags"},
			}}}, bson.D{
				{Key: "$unwind", Value: "$tags"},
			},
		)
	}

	if param.AggregationOpt.Author {
		// set lookup user collection and unwind result
		pipeline = append(pipeline,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.UserTableName},
				{Key: "localField", Value: "authorId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "author"},
			}}},
			bson.D{{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$author"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}}},
		)
	}

	if param.AggregationOpt.Favorite {
		// set lookup favorite collection and addfield size favorite article
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

	// get field selection columns
	articleGroup := bson.D{}
	if len(articleColumns) > 0 {
		for _, column := range articleColumns {
			articleGroup = append(articleGroup, bson.E{Key: column, Value: "$" + column})
		}
	} else {
		article := model.NewArticle()
		for _, column := range article.AllField() {
			articleGroup = append(articleGroup, bson.E{Key: column, Value: "$" + column})
		}
	}

	pipeline = append(pipeline,
		// grouping all result
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "favorite", Value: bson.D{{Key: "$first", Value: "$userFavoritesCount"}}},
			{Key: "author", Value: bson.D{{Key: "$first", Value: "$author"}}},
			{Key: "article", Value: bson.D{{Key: "$first", Value: articleGroup}}},
			{Key: "tags", Value: bson.D{{Key: "$push", Value: bson.D{
				{Key: "_id", Value: "$tags._id"},
				{Key: "name", Value: "$tags.name"},
			}}}},
		}}},
		// set projection by grouping
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "favorite", Value: 1},
			{Key: "author", Value: 1},
			{Key: "article", Value: 1},
			{Key: "tags", Value: 1},
		}}},
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
