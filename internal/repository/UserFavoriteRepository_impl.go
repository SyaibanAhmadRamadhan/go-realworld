package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

type userFavoriteRepositoryImpl struct {
	db *mongo.Database
}

func NewUserFavoriteRepositoryImpl(db *mongo.Database) domain.UserFavoriteRepository {
	return &userFavoriteRepositoryImpl{db: db}
}

func (u *userFavoriteRepositoryImpl) FindAllByUserID(ctx context.Context, param domain.FindAllUserFavoriteParam) (
	res domain.FindAllArticleResult, err error) {
	articleProjection := bson.D{}
	if param.ArticleFields != nil {
		for _, field := range param.ArticleFields {
			articleProjection = append(articleProjection, bson.E{Key: field, Value: "$" + field})
		}
	} else {
		for _, field := range model.NewArticle().AllField() {
			articleProjection = append(articleProjection, bson.E{Key: field, Value: "$" + field})
		}
	}

	pipelineArticleTableName := bson.A{}
	if param.WithTag {
		pipelineArticleTableName = append(pipelineArticleTableName,
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.ArticleTagTableName},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "articleID"},
				{Key: "as", Value: "article_tag"},
			}}},
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: model.TagTableName},
				{Key: "localField", Value: "article_tag.tagID"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "tags"},
			}}},
			bson.D{{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$tags"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}}},
		)
	}
	pipelineArticleTableName = append(pipelineArticleTableName,
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "tags", Value: bson.D{{Key: "$push", Value: "$tags"}}},
			{Key: "articles", Value: bson.D{{Key: "$first", Value: articleProjection}}},
		}}},
	)

	pipeline := mongo.Pipeline{}
	pipeline = append(pipeline,
		bson.D{{Key: "$match", Value: bson.D{
			{Key: "userID", Value: param.UserID},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: model.ArticleTableName},
			{Key: "localField", Value: "articleID"},
			{Key: "foreignField", Value: "_id"},
			{Key: "pipeline", Value: pipelineArticleTableName},
			{Key: "as", Value: "result"},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "tags", Value: "$result.tags"},
			{Key: "article", Value: "$result.articles"},
			{Key: "_id", Value: 0},
		}}},
		bson.D{{Key: "$unwind", Value: "$article"}},
		bson.D{{Key: "$unwind", Value: "$tags"}},
	)

	curTotal, err := u.db.Collection(model.UserFavoriteTableName).Aggregate(ctx, append(pipeline, bson.D{bson.E{Key: "$count", Value: "total"}}))
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
			return res, domain.ErrInvalidTotalType
		}
		res.Total = int64(total)
	}

	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: param.Pagination.Limit}},
		bson.D{{Key: "$skip", Value: param.Pagination.Offset}},
	)
	cur, err := u.db.Collection(model.UserFavoriteTableName).Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	err = cur.All(ctx, &res.Articles)
	return
}

func (u *userFavoriteRepositoryImpl) UpSertByUserID(ctx context.Context, userFavorite model.UserFavorite) (err error) {
	if userFavorite.UserID == "" {
		return domain.ErrIDParamIsEmpty
	}

	filter := bson.D{
		{Key: "userID", Value: userFavorite.UserID},
		{Key: "articleID", Value: userFavorite.ArticleID},
	}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "articleID", Value: userFavorite.ArticleID},
	}}}

	opts := options.Update().SetUpsert(true)
	_, err = u.db.Collection(model.UserFavoriteTableName).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return
	}
	return
}

func (u *userFavoriteRepositoryImpl) DeleteOneByUserID(ctx context.Context, userID string, articleID string) (err error) {
	filter := bson.D{
		{Key: "userID", Value: userID},
		{Key: "articleID", Value: articleID},
	}

	res, err := u.db.Collection(model.UserFavoriteTableName).DeleteOne(ctx, filter)
	if err != nil {
		return
	}

	if res.DeletedCount == 0 {
		err = domain.ErrDelDataNotFound
	}

	return
}
