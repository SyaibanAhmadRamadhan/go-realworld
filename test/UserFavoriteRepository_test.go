package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

var userFavorites []model.UserFavorite

func UserFavoriteRepository_UpSertByUserID(t *testing.T) {
	for _, user := range users {
		for _, article := range articles {
			userFavorite := model.UserFavorite{
				UserID:    user.ID,
				ArticleID: article.ID,
			}
			err := userFavoriteRepository.UpSertByUserID(context.Background(), userFavorite)
			assert.NoError(t, err)
		}
	}
}

func UserFavoriteRepository_FindAllArticleByUserID(t *testing.T) {
	for _, user := range users {
		res, err := userFavoriteRepository.FindAllArticleByUserID(context.Background(), repository.ParamFindAllArticleByUserID{
			WithTag: true,
			Orders: gdb.OrderByParams{
				{Column: "_id", IsAscending: true},
			},
			Pagination: gdb.PaginationParam{
				Limit:  2,
				Offset: 0,
			},
			UserID:        user.ID,
			ArticleFields: []string{"slug"},
		})
		assert.NoError(t, err)
		assert.Equal(t, 2, len(res.Articles))
	}
}
