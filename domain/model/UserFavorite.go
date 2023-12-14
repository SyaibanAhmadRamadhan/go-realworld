package model

type UserFavorite struct {
	UserID    string   `bson:"userID"`
	ArticleID []string `bson:"articleID"`
}

func (u *UserFavorite) TableName() string {
	return "user_favorite"
}
