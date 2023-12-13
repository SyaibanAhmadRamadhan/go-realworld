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

	"realworld-go/domain/model"
	"realworld-go/domain/repository"
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

	// tag collection
	tag := model.Tag{}
	err := mongodb.CreateCollection(ctx, tag.TableName())
	gcommon.PanicIfError(err)

	// article collection
	article := model.Article{}
	err = mongodb.CreateCollection(ctx, article.TableName())
	gcommon.PanicIfError(err)

	// articleTag collection
	articleTag := model.ArticleTag{}
	err = mongodb.CreateCollection(ctx, articleTag.TableName())
	gcommon.PanicIfError(err)

	fmt.Println("finished created collection")
}

var tagRepository repository.TagRepository
var articleRepository repository.ArticleRepository
var articleTagRepository repository.ArticleTagRepository

func initRepository() {
	tagRepository = repositoryimpl.NewTagRepositoryImpl(mongodb)
	articleRepository = repositoryimpl.NewArticleRepositoryImpl(mongodb)
	articleTagRepository = repositoryimpl.NewArticleTagRepositoryImpl(mongodb)
}

func TestRun(t *testing.T) {
	t.Run("TagRepository", func(t *testing.T) {
		t.Run("Create", TagRepository_Create)
		t.Run("FindByID", TagRepository_FindByID)
		t.Run("FindAllByIDS", TagRepository_FindAllByIDS)
		t.Run("UpdateByID", TagRepository_UpdateByID)
		t.Run("DeleteByID", TagRepository_DeleteByID)
	})

	t.Run("ArticleRepository", func(t *testing.T) {
		t.Run("Create", ArticleRepository_Create)
		t.Run("FindByID", ArticleRepository_FindById)
		t.Run("FindAllByIDS", ArticleRepository_FindAllByIDS)
		t.Run("UpdateByID", ArticleRepository_UpdateByID)
		t.Run("DeleteByID", ArticleRepository_DeleteByID)
	})
	t.Run("ArticleTagRepository", func(t *testing.T) {
		t.Run("UpSert", ArticleTagRepository_UpSert)
		t.Run("FindByTagID", ArticleTagRepository_FindByTagID)
		t.Run("FindByArticleID", ArticleTagRepository_FindByArticleID)
		t.Run("FindTagPopuler", ArticleTagRepository_FindTagPopuler)
	})
}
