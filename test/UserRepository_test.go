package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/garray"
	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
)

var users []model.User
var imageDefault = "https://realworld-temp-api.herokuapp.com/images/smiley-cyrus.jpeg"

func UserRepository_Create(t *testing.T) {

	for i := 0; i < 5; i++ {
		user := model.User{
			ID:       gcommon.NewUlid(),
			Email:    gofakeit.Email(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, false, false, 10),
			Demo:     gcommon.RandomFromArray([]bool{true, false}),
			Image:    gcommon.RandomFromArray([]string{imageDefault, gofakeit.ImageURL(640, 480)}),
		}
		fmt.Println(user.Demo)
		err := userRepository.Create(context.Background(), user)
		assert.NoError(t, err)
		users = append(users, user)
	}
}

func UserRepository_FindByOneColumn(t *testing.T) {
	var param []struct {
		filter         []repository.FindByOneColumnParam
		expected       model.User
		columnSelected []string
	}
	for _, user := range users {
		var columnSelected []string
		for i := 0; i < gcommon.RandomFromArray([]int{1, 2, 3}); i++ {
			random := gcommon.RandomFromArray(user.AllField())
			if random == "bio" {
				i--
				continue
			}
			resc, err := garray.AppendUniqueValWithErr(columnSelected, random)
			if err != nil {
				i--
				continue
			}
			columnSelected = resc
		}
		param = append(param, struct {
			filter         []repository.FindByOneColumnParam
			expected       model.User
			columnSelected []string
		}{
			filter: []repository.FindByOneColumnParam{
				{
					Column: "email",
					Value:  user.Email,
				},
				{
					Column: "_id",
					Value:  user.ID,
				},
				{
					Column: "username",
					Value:  user.Username,
				},
			},
			expected:       user,
			columnSelected: columnSelected,
		})
	}

	t.Run("WithoutSelectedColumn", func(t *testing.T) {
		for _, p := range param {
			for _, filter := range p.filter {
				user, err := userRepository.FindByOneColumn(context.Background(), filter)
				assert.NoError(t, err)
				assert.Equal(t, p.expected, user)
			}
		}
	})

	t.Run("WithSelectedColumn", func(t *testing.T) {
		for _, p := range param {
			for _, filter := range p.filter {
				user, err := userRepository.FindByOneColumn(context.Background(), filter, p.columnSelected...)
				assert.NoError(t, err)
				assert.NotEqual(t, p.expected, user)

				values := user.GetValuesByColums(p.columnSelected...)
				for i, column := range p.columnSelected {
					assert.Equal(t, p.expected.GetValuesByColums(column)[0], values[i])
				}
			}
		}
	})
}

func UserRepository_UpdateByID(t *testing.T) {
	var newUsers []model.User
	for _, user := range users {
		bio := gcommon.RandomFromArray([]string{gofakeit.Sentence(10), gofakeit.Sentence(10)})
		user.Username = gofakeit.Username()
		user.Password = gofakeit.Password(true, true, true, false, false, 10)
		user.Image = gofakeit.ImageURL(640, 480)
		user.Bio = &bio
		user.Demo = gcommon.RandomFromArray([]bool{true, false})
		err := userRepository.UpdateByID(context.Background(), user, user.AllField())
		assert.NoError(t, err)

		newUsers = append(newUsers, user)
	}

	users = newUsers
	for _, user := range users {
		res, err := userRepository.FindByOneColumn(context.Background(), repository.FindByOneColumnParam{
			Column: "_id",
			Value:  user.ID,
		})
		assert.NoError(t, err)
		assert.Equal(t, user, res)
		assert.NotNil(t, res.Bio)
	}
}
