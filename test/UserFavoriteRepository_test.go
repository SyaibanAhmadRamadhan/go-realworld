package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

var userFavorites []model.UserFavorite

func UserFavoriteRepository_UpSertByUserID(t *testing.T) {
	for _, user := range users {
		for _, article := range articles {
			userFavorite := model.UserFavorite{
				UserId:    user.Id,
				ArticleId: article.Id,
			}
			err := userFavoriteRepository.UpSertByUserId(context.Background(), userFavorite)
			assert.NoError(t, err)
		}
	}
}

func UserFavoriteRepository_FindAllArticleByUserID(t *testing.T) {
	for _, user := range users {
		res, err := userFavoriteRepository.FindAllByUserId(context.Background(), domain.FindAllUserFavoriteParam{
			WithTag: true,
			Orders: gdb.OrderByParams{
				{Column: "_id", IsAscending: true},
			},
			Pagination: gdb.PaginationParam{
				Limit:  2,
				Offset: 0,
			},
			UserId:        user.Id,
			ArticleFields: []string{"slug"},
		})
		assert.NoError(t, err)
		assert.Equal(t, 2, len(res.Articles))
	}
}
