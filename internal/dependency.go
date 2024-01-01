package internal

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gmongodb"
	"github.com/SyaibanAhmadRamadhan/gocatch/gvalidation"

	"realworld-go/domain"
	"realworld-go/infra"
	"realworld-go/internal/repository"
	"realworld-go/internal/usecase"
)

type DependencyConfig struct {
	Mongo *infra.MongoDB
	Minio *infra.Minio
}

type Dependency struct {
	ArticleUsecase domain.ArticleUsecase
	CommentUsecase domain.CommentUsecase
	TagUsecase     domain.TagUsecase
}

func DependencyMongodb(dependencyConfig DependencyConfig) *Dependency {
	validation := gvalidation.New()
	validation.SetEnTranslation()

	articleRepo := repository.NewArticleRepositoryImpl(dependencyConfig.Mongo.Database)
	articleTagRepo := repository.NewArticleTagRepositoryImpl(dependencyConfig.Mongo.Database)
	commentRepo := repository.NewCommentRepositoryImpl(dependencyConfig.Mongo.Database)
	tagRepo := repository.NewTagRepositoryImpl(dependencyConfig.Mongo.Database)
	userRepo := repository.NewUserRepositoryImpl(dependencyConfig.Mongo.Database)
	// userFavoriteRepo := repository.NewUserFavoriteRepositoryImpl(mdb)
	txRepo := gmongodb.NewTxMongodb(dependencyConfig.Mongo.Client)

	articleUsecase := usecase.NewArticleUsecaseImpl(articleRepo, articleTagRepo, commentRepo, tagRepo, txRepo, validation)
	commentUsecase := usecase.NewCommentUsecaseImpl(commentRepo, userRepo, validation)
	tagUsecase := usecase.NewTagUsecaseImpl(tagRepo)
	return &Dependency{
		ArticleUsecase: articleUsecase,
		CommentUsecase: commentUsecase,
		TagUsecase:     tagUsecase,
	}
}
