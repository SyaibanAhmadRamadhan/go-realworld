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

var articleTags [][]model.ArticleTag

func ArticleTagRepository_ReplaceAll(t *testing.T) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	t.Run("Insert", func(t *testing.T) {
		for _, article := range articles {
			var articleTag []model.ArticleTag
			for i := 0; i < gcommon.RandomFromArray([]int{4, 5, 6}); i++ {
				articleTag = garray.AppendUniqueVal(articleTag, model.ArticleTag{
					ArticleID: article.ID,
					TagID:     gcommon.RandomFromArray(tagIDs),
				})
			}

			articleTags = append(articleTags, articleTag)

			err := articleTagRepository.ReplaceAll(context.Background(), articleTag)
			assert.NoError(t, err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		articleTags = nil
		for _, article := range articles {
			var articleTag []model.ArticleTag
			for i := 0; i < gcommon.RandomFromArray([]int{1, 2, 3}); i++ {
				articleTag = garray.AppendUniqueVal(articleTag, model.ArticleTag{
					ArticleID: article.ID,
					TagID:     gcommon.RandomFromArray(tagIDs),
				})
			}

			articleTags = append(articleTags, articleTag)

			err := articleTagRepository.ReplaceAll(context.Background(), articleTag)
			assert.NoError(t, err)
		}
	})
}

func ArticleTagRepository_FindAllDetail(t *testing.T) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = garray.AppendUniqueVal(tagIDs, tag.ID)
	}
	res, err := articleTagRepository.FindAllDetail(context.Background(), repository.ParamFindAllDetailAT{
		TagIDs: tagIDs,
		OrderBy: repository.OrderBy{
			Column:      "slug",
			IsAscending: true,
		},
		Pagination: repository.PaginationParam{
			Limit:  5,
			Offset: 0,
		},
	}, "slug")
	assert.NoError(t, err)
	assert.Equal(t, len(articles), int(res.Total))
}

func ArticleTagRepository_FindOneByArticleID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		for _, articleTag := range articleTags {
			res, err := articleTagRepository.FindOneByArticleID(context.Background(), articleTag[0].ArticleID)
			assert.NoError(t, err)
			assert.Equal(t, articleTag[0].ArticleID, res.Article.ID)
			assert.Equal(t, len(articleTag), len(res.Tags))
		}
	})

	t.Run("Failed", func(t *testing.T) {
		res, err := articleTagRepository.FindOneByArticleID(context.Background(), "articleTag.ArticleID")
		assert.NoError(t, err)
		assert.Equal(t, res.Article.ID, "")
	})
}

func ArticleTagRepository_FindTagPopuler(t *testing.T) {
	res, err := articleTagRepository.FindTagPopuler(context.Background(), 5)
	assert.NoError(t, err)
	for _, re := range res {
		resArticle, errArticle := articleTagRepository.FindAllDetail(context.Background(), repository.ParamFindAllDetailAT{
			TagIDs: []string{re.TagID},
			OrderBy: repository.OrderBy{
				Column:      "_id",
				IsAscending: true,
			},
			Pagination: repository.PaginationParam{
				Limit:  2,
				Offset: 0,
			},
		})
		assert.NoError(t, errArticle)
		assert.Equal(t, re.Count, resArticle.Total)
	}
}

//	func ArticleTagRepository_FindByTagID(t *testing.T) {
//		var articleTotal int
//
//		for _, tag := range tags {
//			var limit int64 = 5
//			var offset int64 = 0
//
//			for {
//				res, total, err := articleTagRepository.FindByTagID(context.Background(), tag.ID, repository.PaginationParam{
//					Limit:  limit,
//					Offset: offset,
//				})
//				assert.NoError(t, err)
//
//				if offset == 0 {
//					articleTotal += int(total)
//				}
//				if total > limit {
//					if len(res) == 0 {
//
//					} else if total%limit == 0 {
//						assert.Equal(t, 5, len(res))
//					} else {
//						if len(res) < int(limit) {
//							final := total % limit
//							if final > total {
//								final = final - total
//							}
//							assert.Equal(t, int(final), len(res))
//						} else {
//							assert.Equal(t, int(limit), len(res))
//						}
//					}
//				}
//
//				if len(res) < int(limit) {
//					break
//				}
//
//				offset += limit
//			}
//		}
//
//		assert.NotEqual(t, len(articles), articleTotal)
//	}
//
