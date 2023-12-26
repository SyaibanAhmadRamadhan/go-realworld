package model

const UserFavoriteTableName = "user_favorite"

type UserFavorite struct {
	UserId    string `bson:"userId"`
	ArticleId string `bson:"articleId"`
}
