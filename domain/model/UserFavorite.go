package model

const UserFavoriteTableName = "user_favorite"

type UserFavorite struct {
	UserID    string `bson:"userID"`
	ArticleID string `bson:"articleID"`
}
