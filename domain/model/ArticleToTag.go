package model

type ArticleToTag struct {
	ArticleID int `bson:"articleID"`
	TagID     int `bson:"tagID"`
}
