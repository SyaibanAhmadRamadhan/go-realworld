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
	"realworld-go/internal/repository"
)

var articles []model.Article

func ArticleRepository_Create(t *testing.T) {
	for i := 0; i < 50; i++ {
		createdAt := gofakeit.Date()
		createdAt = gtime.NormalizeTimeUnit(createdAt, gtime.Milliseconds)
		var userIDs []string
		for _, user := range users {
			userIDs = append(userIDs, user.Id)
		}
		article := model.Article{
			Id:          gcommon.NewUlid() + "_article",
			AuthorId:    gcommon.RandomFromArray(userIDs),
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
			res, err := articleRepository.FindOneByOneColumn(context.Background(), domain.FindOneByIdArticleParam{
				ArticleId: article.Id,
				AggregationOpt: domain.FindArticleOpt{
					Tag:      false,
					Favorite: true,
				},
			}, "slug")
			assert.NoError(t, err)
			assert.Equal(t, article.Slug, res.Article.Slug)
			assert.NotEqual(t, article, res.Article)

			res1, err := articleRepository.FindOneByOneColumn(context.Background(), domain.FindOneByIdArticleParam{
				ArticleId:      article.Id,
				AggregationOpt: domain.FindArticleOpt{},
			})
			assert.NoError(t, err)
			assert.Equal(t, res1.Article, article)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		_, err := articleRepository.FindOneByOneColumn(context.Background(), domain.FindOneByIdArticleParam{
			ArticleId: "article.Id",
		})
		assert.Equal(t, repository.ErrDataNotFound, err)
	})
}

func ArticleRepository_FindAllPaginate(t *testing.T) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = garray.AppendUniqueVal(tagIDs, tag.Id)
	}
	res, err := articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
		TagIds: tagIDs,
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
			err := articleRepository.UpdateById(context.Background(), articleUpdate.Article, columns)
			assert.NoError(t, err)

			res, err := articleRepository.FindOneByOneColumn(context.Background(), domain.FindOneByIdArticleParam{
				ArticleId: articleUpdate.Article.Id,
			})
			assert.NoError(t, err)
			assert.Equal(t, articleUpdate.Expected, res.Article.Slug)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		err := articleRepository.UpdateById(context.Background(), model.Article{
			Id:   "random",
			Slug: gofakeit.Slogan(),
		}, columns)
		assert.Equal(t, repository.ErrUpdateDataNotFound, err)
	})
}

func ArticleRepository_DeleteByID(t *testing.T) {
	var articleDeleteds []model.Article
	var ids []string
	for i := 0; i < 5; i++ {
		article := model.Article{
			Id:   gcommon.NewUlid(),
			Slug: gofakeit.Slogan(),
		}
		articleDeleteds = append(articleDeleteds, article)
		ids = append(ids, article.Id)

		err := articleRepository.Create(context.Background(), article)
		gcommon.PanicIfError(err)
	}

	res, err := articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
		TagIds: nil,
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
			err = articleRepository.DeleteById(context.Background(), articleDeleted)
			assert.NoError(t, err)

			_, err = articleRepository.FindOneByOneColumn(context.Background(), domain.FindOneByIdArticleParam{
				ArticleId: articleDeleted.Id,
			})
			assert.Equal(t, repository.ErrDataNotFound, err)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		err = tagRepository.DeleteById(context.Background(), model.Tag{
			Id: "random",
		})
		assert.Equal(t, repository.ErrDelDataNotFound, err)
	})

	res, err = articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
		TagIds: nil,
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
