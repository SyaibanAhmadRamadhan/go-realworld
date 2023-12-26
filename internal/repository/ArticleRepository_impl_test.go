package repository

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/gtime"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

func TestArticleRepositoryImpl_Create(t *testing.T) {
	ctx := context.Background()
	opts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, opts)

	var articles []model.Article
	for i := 0; i < 10; i++ {
		articles = append(articles, model.Article{
			Id:          gcommon.NewUlid(),
			AuthorId:    gcommon.NewUlid(),
			Slug:        gofakeit.Username(),
			Title:       gofakeit.Sentence(5),
			Description: gofakeit.Paragraph(2, 2, 10, "\n"),
			Body:        gofakeit.Paragraph(5, 4, 10, "\n"),
			CreatedAt:   gtime.NormalizeTimeUnit(gofakeit.Date(), gtime.Milliseconds),
			UpdatedAt:   gtime.NormalizeTimeUnit(gofakeit.Date(), gtime.Milliseconds),
		})
	}

	mt.Run("success", func(t *mtest.T) {
		articleRepo := NewArticleRepositoryImpl(t.DB)
		for _, article := range articles {
			t.AddMockResponses(mtest.CreateSuccessResponse())
			err := articleRepo.Create(ctx, article)
			assert.NoError(t, err)
		}
	})
}

func TestArticleRepositoryImpl_FindOneByID(t *testing.T) {
	ctx := context.Background()
	opts := mtest.NewOptions().ClientType(mtest.Mock)
	mt := mtest.New(t, opts)

	var articles []model.Article
	for i := 0; i < 10; i++ {
		articles = append(articles, model.Article{
			Id:          gcommon.NewUlid(),
			AuthorId:    gcommon.NewUlid(),
			Slug:        gofakeit.Username(),
			Title:       gofakeit.Sentence(5),
			Description: gofakeit.Paragraph(2, 2, 10, "\n"),
			Body:        gofakeit.Paragraph(5, 4, 10, "\n"),
			CreatedAt:   gtime.NormalizeTimeUnit(gofakeit.Date(), gtime.Milliseconds),
			UpdatedAt:   gtime.NormalizeTimeUnit(gofakeit.Date(), gtime.Milliseconds),
		})
	}

	mt.Run("success", func(mt *mtest.T) {
		articleRepo := NewArticleRepositoryImpl(mt.DB)
		for _, article := range articles {
			mt.AddMockResponses(mtest.CreateCursorResponse(0, "dbname.collname", mtest.FirstBatch, bson.D{
				{Key: "article", Value: bson.D{
					{Key: "_id", Value: article.Id},
					{Key: article.FieldAuthorId(), Value: article.AuthorId},
					{Key: article.FieldSlug(), Value: article.Slug},
					{Key: article.FieldTitle(), Value: article.Title},
					{Key: article.FieldDescription(), Value: article.Description},
					{Key: article.FieldBody(), Value: article.Body},
					{Key: article.FieldCreatedAt(), Value: article.CreatedAt},
					{Key: article.FieldUpdatedAt(), Value: article.UpdatedAt},
				}},
			}))
			res, err := articleRepo.FindOneById(ctx, domain.FindOneByIdArticleParam{
				ArticleId: article.Id,
				AggregationOpt: domain.FindArticleOpt{
					Tag:      true,
					Favorite: true,
					Author:   false,
				},
			})
			assert.NoError(mt, err)
			assert.Equal(t, article, res.Article)
		}
	})
}
