package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gmongodb"
	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain"
	"realworld-go/domain/model"
	repositoryimpl "realworld-go/internal/repository"
)

var mongoClient *mongo.Client
var mongodb *mongo.Database

const mongoNameDB = "realworld-mongo"

func TestMain(m *testing.M) {
	dockerTest := ginfra.InitDockerTest()
	defer dockerTest.CleanUp()

	// mongodb dockertest
	mongodbDockerTestConf := gmongodb.MongoDockerTestConf{}
	dockerTest.NewContainer(mongodbDockerTestConf.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
		time.Sleep(5 * time.Second)
		var err error
		mongoClient, err = mongodbDockerTestConf.ConnectClient(res)
		if err != nil {
			return err
		}
		return nil
	})

	// init collection
	createCollection()

	// init repository layer
	initRepository()

	// run test
	m.Run()
}

func createCollection() {
	ctx := context.Background()
	mongodb = mongoClient.Database(mongoNameDB)

	err := mongodb.CreateCollection(ctx, model.TagTableName)
	gcommon.PanicIfError(err)

	err = mongodb.CreateCollection(ctx, model.ArticleTableName)
	gcommon.PanicIfError(err)

	err = mongodb.CreateCollection(ctx, model.ArticleTagTableName)
	gcommon.PanicIfError(err)

	fmt.Println("finished created collection")
}

var tagRepository domain.TagRepository
var userRepository domain.UserRepository
var articleRepository domain.ArticleRepository
var articleTagRepository domain.ArticleTagRepository
var userFavoriteRepository domain.UserFavoriteRepository

func initRepository() {
	tagRepository = repositoryimpl.NewTagRepositoryImpl(mongodb)
	userRepository = repositoryimpl.NewUserRepositoryImpl(mongodb)
	articleRepository = repositoryimpl.NewArticleRepositoryImpl(mongodb)
	articleTagRepository = repositoryimpl.NewArticleTagRepositoryImpl(mongodb)
	userFavoriteRepository = repositoryimpl.NewUserFavoriteRepositoryImpl(mongodb)
}

func TestRun(t *testing.T) {
	t.Run("TagRepository", func(t *testing.T) {
		t.Run("Create", TagRepository_Create)
		t.Run("FindOneByID", TagRepository_FindByID)
		t.Run("FindAllByIDs", TagRepository_FindAllByIDS)
		t.Run("UpSertByID", TagRepository_UpdateByID)
		t.Run("DeleteByID", TagRepository_DeleteByID)
	})

	t.Run("UserRepository", func(t *testing.T) {
		t.Run("Create", UserRepository_Create)
		t.Run("FindByOneColumn", UserRepository_FindByOneColumn)
		t.Run("UpSertByID", UserRepository_UpdateByID)
	})

	t.Run("ArticleRepository", func(t *testing.T) {
		t.Run("Create", ArticleRepository_Create)
	})

	t.Run("ArticleTagRepository", func(t *testing.T) {
		t.Run("ReplaceAll", ArticleTagRepository_ReplaceAll)
	})

	t.Run("TagRepository", func(t *testing.T) {
		t.Run("FindTagPopuler", TagRepository_FindTagPopuler)
	})

	t.Run("UserFavoriteRepository", func(t *testing.T) {
		t.Run("UpSertByUserID", UserFavoriteRepository_UpSertByUserID)
		t.Run("FindAllByUserID", UserFavoriteRepository_FindAllArticleByUserID)
	})

	t.Run("ArticleRepository", func(t *testing.T) {
		t.Run("FindOneByID", ArticleRepository_FindOneByID)
		t.Run("FindAllPaginate", ArticleRepository_FindAllPaginate)
		t.Run("UpSertByID", ArticleRepository_UpdateByID)
		t.Run("DeleteByID", ArticleRepository_DeleteByID)
	})
}
