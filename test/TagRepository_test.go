package test

import (
	"context"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain"
	"realworld-go/domain/model"
)

var tags []model.Tag

// func TagRepository_Create(t *testing.T) {
// 	for i := 0; i < 20; i++ {
// 		tag := model.Tag{
// 			Id:   gcommon.NewUlid(),
// 			Name: gofakeit.Gamertag(),
// 		}
// 		tags = append(tags, tag)
//
// 		err := tagRepository.UpSertMany(context.Background(), tag)
// 		gcommon.PanicIfError(err)
// 	}
// }

// func TagRepository_FindByID(t *testing.T) {
// 	t.Run("Success", func(t *testing.T) {
// 		for _, tag := range tags {
// 			res, err := tagRepository.FindById(context.Background(), tag.Id)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tag.Name, res.Name)
// 		}
// 	})
//
// 	t.Run("Failed", func(t *testing.T) {
// 		_, err := tagRepository.FindByID(context.Background(), "tag.Id")
// 		assert.Equal(t, domain.ErrDataNotFound, err)
// 	})
// }
//
// func TagRepository_FindAllByIDS(t *testing.T) {
// 	var ids []string
// 	for _, tag := range tags {
// 		ids = append(ids, tag.Id)
// 	}
// 	res, err := tagRepository.FindAllByIDS(context.Background(), ids)
// 	assert.NoError(t, err)
// 	assert.Equal(t, tags, res)
// }

func TagRepository_FindTagPopuler(t *testing.T) {
	res, err := tagRepository.FindTagPopuler(context.Background(), 5)
	assert.NoError(t, err)
	for _, re := range res {
		resArticle, errArticle := articleRepository.FindAllPaginate(context.Background(), domain.FindAllPaginateArticleParam{
			TagIds: []string{re.TagId},
			Orders: gdb.OrderByParams{
				{Column: "slug", IsAscending: true},
			},
			Pagination: gdb.PaginationParam{
				Limit:  2,
				Offset: 0,
			},
			AggregationOpt: domain.FindArticleOpt{
				Tag:      true,
				Favorite: false,
			},
		})
		assert.NoError(t, errArticle)
		assert.Equal(t, re.Count, resArticle.Total)
	}
}

// func TagRepository_UpdateByID(t *testing.T) {
// 	var tagUpdates []struct {
// 		Tag      model.Tag
// 		Expected string
// 	}
//
// 	for _, tag := range tags {
// 		tagName := gofakeit.AppName()
// 		tag.Name = tagName
// 		tagUpdate := struct {
// 			Tag      model.Tag
// 			Expected string
// 		}{
// 			Tag:      tag,
// 			Expected: tagName,
// 		}
// 		tagUpdates = append(tagUpdates, tagUpdate)
// 	}
//
// 	columns := []string{
// 		tags[0].FieldName(),
// 	}
//
// 	t.Run("success", func(t *testing.T) {
// 		for _, tagUpdate := range tagUpdates {
// 			err := tagRepository.UpdateByID(context.Background(), tagUpdate.Tag, columns)
// 			assert.NoError(t, err)
//
// 			res, err := tagRepository.FindByID(context.Background(), tagUpdate.Tag.Id)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tagUpdate.Expected, res.Name)
// 		}
// 	})
//
// 	t.Run("Failed", func(t *testing.T) {
// 		err := tagRepository.UpdateByID(context.Background(), model.Tag{
// 			Id: "random",
// 		}, columns)
// 		assert.Equal(t, domain.ErrUpdateDataNotFound, err)
// 	})
// }
//
// func TagRepository_DeleteByID(t *testing.T) {
// 	var tagDeleteds []model.Tag
// 	var ids []string
// 	for i := 0; i < 5; i++ {
// 		tag := model.Tag{
// 			Id:   gofakeit.UUID(),
// 			Name: gofakeit.Gamertag(),
// 		}
// 		tagDeleteds = append(tagDeleteds, tag)
// 		ids = append(ids, tag.Id)
//
// 		err := tagRepository.Create(context.Background(), tag)
// 		gcommon.PanicIfError(err)
// 	}
//
// 	res, err := tagRepository.FindAllByIDS(context.Background(), ids)
// 	assert.NoError(t, err)
// 	assert.Equal(t, len(tagDeleteds), len(res))
//
// 	t.Run("Success", func(t *testing.T) {
// 		for _, tagDeleted := range tagDeleteds {
// 			err = tagRepository.DeleteByID(context.Background(), tagDeleted)
// 			assert.NoError(t, err)
//
// 			_, err = tagRepository.FindByID(context.Background(), tagDeleted.Id)
// 			assert.Equal(t, domain.ErrDataNotFound, err)
// 		}
// 	})
//
// 	t.Run("Failed", func(t *testing.T) {
// 		err = tagRepository.DeleteById(context.Background(), model.Tag{
// 			Id: "random",
// 		})
// 		assert.Equal(t, domain.ErrDelDataNotFound, err)
// 	})
//
// 	res, err = tagRepository.FindAllByIDS(context.Background(), ids)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 0, len(res))
// }
