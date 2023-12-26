package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/garray"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain/model"
)

var articleTags [][]model.ArticleTag

func ArticleTagRepository_ReplaceAll(t *testing.T) {
	var tagIDs []string
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.Id)
	}

	t.Run("Insert", func(t *testing.T) {
		for _, article := range articles {
			var articleTag []model.ArticleTag
			for i := 0; i < gcommon.RandomFromArray([]int{4, 5, 6}); i++ {
				articleTag = garray.AppendUniqueVal(articleTag, model.ArticleTag{
					ArticleId: article.Id,
					TagId:     gcommon.RandomFromArray(tagIDs),
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
					ArticleId: article.Id,
					TagId:     gcommon.RandomFromArray(tagIDs),
				})
			}

			articleTags = append(articleTags, articleTag)

			err := articleTagRepository.ReplaceAll(context.Background(), articleTag)
			assert.NoError(t, err)
		}
	})
}

//	func ArticleTagRepository_FindByTagID(t *testing.T) {
//		var articleTotal int
//
//		for _, tag := range tags {
//			var limit int64 = 5
//			var offset int64 = 0
//
//			for {
//				res, total, err := articleTagRepository.FindByTagID(context.Background(), tag.Id, repository.PaginationParam{
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
