package model

type UserFavorite struct {
	UserID    int `bson:"userID"`
	ArticleID int `bson:"articleID"`
}
