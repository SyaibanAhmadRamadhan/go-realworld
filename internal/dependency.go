package internal

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gmongodb"
	"github.com/SyaibanAhmadRamadhan/gocatch/gvalidation"
	"go.mongodb.org/mongo-driver/mongo"

	"realworld-go/domain"
	"realworld-go/internal/repository"
	"realworld-go/internal/usecase"
)

type Dependency struct {
	ArticleUsecase domain.ArticleUsecase
	CommentUsecase domain.CommentUsecase
	TagUsecase     domain.TagUsecase
}

func DependencyMongodb(mdb *mongo.Database, mclient *mongo.Client) *Dependency {
	validation := gvalidation.New()
	validation.SetEnTranslation()

	articleRepo := repository.NewArticleRepositoryImpl(mdb)
	articleTagRepo := repository.NewArticleTagRepositoryImpl(mdb)
	commentRepo := repository.NewCommentRepositoryImpl(mdb)
	tagRepo := repository.NewTagRepositoryImpl(mdb)
	userRepo := repository.NewUserRepositoryImpl(mdb)
	// userFavoriteRepo := repository.NewUserFavoriteRepositoryImpl(mdb)
	txRepo := gmongodb.NewTxMongodb(mclient)

	articleUsecase := usecase.NewArticleUsecaseImpl(articleRepo, articleTagRepo, commentRepo, tagRepo, txRepo, validation)
	commentUsecase := usecase.NewCommentUsecaseImpl(commentRepo, userRepo, validation)
	tagUsecase := usecase.NewTagUsecaseImpl(tagRepo)
	return &Dependency{
		ArticleUsecase: articleUsecase,
		CommentUsecase: commentUsecase,
		TagUsecase:     tagUsecase,
	}
}
