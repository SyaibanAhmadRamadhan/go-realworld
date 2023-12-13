package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/garray"
	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

var articleTags []model.ArticleTag

func ArticleTagRepository_UpSert(t *testing.T) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	t.Run("Insert", func(t *testing.T) {
		for _, article := range articles {
			var tag []string
			for i := 0; i < gcommon.RandomFromArray([]int{4, 5, 6}); i++ {
				tag = garray.AppendUniqueVal(tag, gcommon.RandomFromArray(tagIDs))
			}
			articleTag := model.ArticleTag{
				ArticleID: article.ID,
				TagIDs:    tag,
			}
			articleTags = append(articleTags, articleTag)

			err := articleTagRepository.UpSert(context.Background(), articleTag)
			assert.NoError(t, err)

			res, err := articleTagRepository.FindByArticleID(context.Background(), articleTag.ArticleID)
			assert.NoError(t, err)
			assert.Equal(t, tag, res.TagIDs)
		}
	})

	t.Run("Update", func(t *testing.T) {
		articleTags = nil
		for _, article := range articles {
			var tag []string
			for i := 0; i < gcommon.RandomFromArray([]int{1, 2, 3}); i++ {
				tag = garray.AppendUniqueVal(tag, gcommon.RandomFromArray(tagIDs))
			}

			articleTag := model.ArticleTag{
				ArticleID: article.ID,
				TagIDs:    tag,
			}
			articleTags = append(articleTags, articleTag)

			err := articleTagRepository.UpSert(context.Background(), articleTag)
			assert.NoError(t, err)

			res, err := articleTagRepository.FindByArticleID(context.Background(), articleTag.ArticleID)
			assert.NoError(t, err)
			assert.Equal(t, tag, res.TagIDs)
		}
	})
}

func ArticleTagRepository_FindByTagID(t *testing.T) {
	var articleTotal int

	for _, tag := range tags {
		var limit int64 = 5
		var offset int64 = 0

		for {
			res, total, err := articleTagRepository.FindByTagID(context.Background(), tag.ID, repository.PaginationParam{
				Limit:  limit,
				Offset: offset,
			})
			assert.NoError(t, err)

			if offset == 0 {
				articleTotal += int(total)
			}
			if total > limit {
				if len(res) == 0 {

				} else if total%limit == 0 {
					assert.Equal(t, 5, len(res))
				} else {
					if len(res) < int(limit) {
						final := total % limit
						if final > total {
							final = final - total
						}
						assert.Equal(t, int(final), len(res))
					} else {
						assert.Equal(t, int(limit), len(res))
					}
				}
			}

			if len(res) < int(limit) {
				break
			}

			offset += limit
		}
	}

	assert.NotEqual(t, len(articles), articleTotal)
}

func ArticleTagRepository_FindByArticleID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		for _, articleTag := range articleTags {
			res, err := articleTagRepository.FindByArticleID(context.Background(), articleTag.ArticleID)
			assert.NoError(t, err)
			assert.Equal(t, articleTag.TagIDs, res.TagIDs)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		res, err := articleTagRepository.FindByArticleID(context.Background(), "articleTag.ArticleID")
		assert.ErrorIs(t, err, repository.ErrDataNotFound)
		assert.Equal(t, "", res.ArticleID)
	})
}

func ArticleTagRepository_FindTagPopuler(t *testing.T) {
	res, err := articleTagRepository.FindTagPopuler(context.Background(), 100)
	assert.NoError(t, err)

	for _, re := range res {
		_, total, err := articleTagRepository.FindByTagID(context.Background(), re.TagIDs, repository.PaginationParam{
			Limit:  2,
			Offset: 0,
		})
		assert.NoError(t, err)
		assert.Equal(t, re.Count, total)
	}
}
