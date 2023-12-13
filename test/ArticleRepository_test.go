package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtime"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

var articles []model.Article

func ArticleRepository_Create(t *testing.T) {
	for i := 0; i < 50; i++ {
		createdAt := gofakeit.Date()
		createdAt = gtime.NormalizeTimeUnit(createdAt, gtime.Milliseconds)
		article := model.Article{
			ID:          gcommon.NewUlid(),
			AuthorID:    gofakeit.Int64(),
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

func ArticleRepository_FindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		for _, article := range articles {
			res, err := articleRepository.FindById(context.Background(), article.ID)
			assert.NoError(t, err)
			assert.Equal(t, article, res)

			res1, err := articleRepository.FindById(context.Background(), article.ID, article.FieldBody())
			assert.NoError(t, err)
			assert.Equal(t, res1.Body, article.Body)
			assert.NotEqual(t, res1.Slug, article.Slug)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		_, err := articleRepository.FindById(context.Background(), "article.ID")
		assert.Equal(t, repository.ErrDataNotFound, err)
	})
}

func ArticleRepository_FindAllByIDS(t *testing.T) {
	var articleSelectedColumns []model.Article
	var ids []string
	for _, article := range articles {
		ids = append(ids, article.ID)
		articleSelectedColumns = append(articleSelectedColumns, model.Article{
			ID:   article.ID,
			Body: article.Body,
			Slug: article.Slug,
		})

	}

	res, err := articleRepository.FindAllByIDS(context.Background(), ids)
	assert.NoError(t, err)
	assert.Equal(t, articles, res)

	var article model.Article
	res, err = articleRepository.FindAllByIDS(context.Background(), ids, article.FieldSlug(), article.FieldBody())
	assert.NoError(t, err)
	assert.Equal(t, articleSelectedColumns, res)
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

			res, err := articleRepository.FindById(context.Background(), articleUpdate.Article.ID)
			assert.NoError(t, err)
			assert.Equal(t, articleUpdate.Expected, res.Slug)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		err := articleRepository.UpdateByID(context.Background(), model.Article{
			ID:   "random",
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
			ID:   gcommon.NewUlid(),
			Slug: gofakeit.Slogan(),
		}
		articleDeleteds = append(articleDeleteds, article)
		ids = append(ids, article.ID)

		err := articleRepository.Create(context.Background(), article)
		gcommon.PanicIfError(err)
	}

	res, err := articleRepository.FindAllByIDS(context.Background(), ids)
	assert.NoError(t, err)
	assert.Equal(t, len(articleDeleteds), len(res))

	t.Run("Success", func(t *testing.T) {
		for _, articleDeleted := range articleDeleteds {
			err = articleRepository.DeleteByID(context.Background(), articleDeleted)
			assert.NoError(t, err)

			_, err = articleRepository.FindById(context.Background(), articleDeleted.ID)
			assert.Equal(t, repository.ErrDataNotFound, err)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		err = tagRepository.DeleteByID(context.Background(), model.Tag{
			ID: "random",
		})
		assert.Equal(t, repository.ErrDelDataNotFound, err)
	})

	res, err = articleRepository.FindAllByIDS(context.Background(), ids)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
}
