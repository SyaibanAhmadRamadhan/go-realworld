package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/garray"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/gtime"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

var articles []model.Article

func ArticleRepository_Create(t *testing.T) {
	for i := 0; i < 50; i++ {
		createdAt := gofakeit.Date()
		createdAt = gtime.NormalizeTimeUnit(createdAt, gtime.Milliseconds)
		var userIDs []string
		for _, user := range users {
			userIDs = append(userIDs, user.ID)
		}
		article := model.Article{
			ID:          gcommon.NewUlid() + "_article",
			AuthorID:    gcommon.RandomFromArray(userIDs),
			Slug:        gofakeit.Username(),
			Title:       gofakeit.Sentence(5),
			Description: gofakeit.Paragraph(2, 2, 10, "\n"),
			Body:        gofakeit.Paragraph(5, 4, 10, "\n"),
			CreatedAt:   createdAt,
			UpdatedAt:   createdAt,
		}
		articles = append(articles, article)

		err := articleRepository.Create(context.Background(), article)
		gcommon.PanicIfError(err)
	}
}

func ArticleRepository_FindOneByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		for _, article := range articles {
			res, err := articleRepository.FindOneByID(context.Background(), domain.FindOneByIDArticleParam{
				ArticleID: article.ID,
				AggregationOpt: domain.FindArticleOpt{
					Tag:      false,
					Favorite: true,
				},
			}, "slug")
			assert.NoError(t, err)
			assert.Equal(t, article.Slug, res.Article.Slug)
			assert.NotEqual(t, article, res.Article)

			res1, err := articleRepository.FindOneByID(context.Background(), domain.FindOneByIDArticleParam{
				ArticleID:      article.ID,
				AggregationOpt: domain.FindArticleOpt{},
			})
			assert.NoError(t, err)
			assert.Equal(t, res1.Article, article)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		_, err := articleRepository.FindOneByID(context.Background(), domain.FindOneByIDArticleParam{
			ArticleID: "article.ID",
		})
		assert.Equal(t, domain.ErrDataNotFound, err)
	})
}

func ArticleRepository_FindAllPaginate(t *testing.T) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = garray.AppendUniqueVal(tagIDs, tag.ID)
	}
	res, err := articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
		TagIDs: tagIDs,
		Orders: gdb.OrderByParams{
			{Column: "slug", IsAscending: true},
			{Column: "asal", IsAscending: true},
		},
		Pagination: gdb.PaginationParam{
			Limit:  5,
			Offset: 0,
		},
		AggregationOpt: domain.FindArticleOpt{
			Tag:      true,
			Favorite: true,
		},
	}, "slug")
	assert.NoError(t, err)
	assert.Equal(t, len(articles), int(res.Total))
}

func ArticleRepository_UpdateByID(t *testing.T) {
	var articleUpdates []struct {
		Article  model.Article
		Expected string
	}

	for _, article := range articles {
		slug := gofakeit.Slogan()
		article.Slug = slug
		articleUpdate := struct {
			Article  model.Article
			Expected string
		}{
			Article:  article,
			Expected: slug,
		}
		articleUpdates = append(articleUpdates, articleUpdate)
	}

	columns := []string{
		articles[0].FieldSlug(),
	}

	t.Run("success", func(t *testing.T) {
		for _, articleUpdate := range articleUpdates {
			err := articleRepository.UpdateByID(context.Background(), articleUpdate.Article, columns)
			assert.NoError(t, err)

			res, err := articleRepository.FindOneByID(context.Background(), domain.FindOneByIDArticleParam{
				ArticleID: articleUpdate.Article.ID,
			})
			assert.NoError(t, err)
			assert.Equal(t, articleUpdate.Expected, res.Article.Slug)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		err := articleRepository.UpdateByID(context.Background(), model.Article{
			ID:   "random",
			Slug: gofakeit.Slogan(),
		}, columns)
		assert.Equal(t, domain.ErrUpdateDataNotFound, err)
	})
}

func ArticleRepository_DeleteByID(t *testing.T) {
	var articleDeleteds []model.Article
	var ids []string
	for i := 0; i < 5; i++ {
		article := model.Article{
			ID:   gcommon.NewUlid(),
			Slug: gofakeit.Slogan(),
		}
		articleDeleteds = append(articleDeleteds, article)
		ids = append(ids, article.ID)

		err := articleRepository.Create(context.Background(), article)
		gcommon.PanicIfError(err)
	}

	res, err := articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
		TagIDs: nil,
		Orders: gdb.OrderByParams{
			{Column: "slug", IsAscending: true},
		},
		Pagination: gdb.PaginationParam{
			Limit:  5,
			Offset: 0,
		},
	}, "slug")
	assert.NoError(t, err)
	assert.Equal(t, len(articles)+5, int(res.Total))

	t.Run("Success", func(t *testing.T) {
		for _, articleDeleted := range articleDeleteds {
			err = articleRepository.DeleteByID(context.Background(), articleDeleted)
			assert.NoError(t, err)

			_, err = articleRepository.FindOneByID(context.Background(), domain.FindOneByIDArticleParam{
				ArticleID: articleDeleted.ID,
			})
			assert.Equal(t, domain.ErrDataNotFound, err)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		err = tagRepository.DeleteByID(context.Background(), model.Tag{
			ID: "random",
		})
		assert.Equal(t, domain.ErrDelDataNotFound, err)
	})

	res, err = articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
		TagIDs: nil,
		Orders: gdb.OrderByParams{
			{Column: "slug", IsAscending: true},
		},
		Pagination: gdb.PaginationParam{
			Limit:  5,
			Offset: 0,
		},
	}, "slug")
	assert.NoError(t, err)
	assert.Equal(t, len(articles), int(res.Total))
}
