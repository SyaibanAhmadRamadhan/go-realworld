package model

const ArticleTagTableName string = "articleTag"

type ArticleTag struct {
	ArticleID string `bson:"articleID"`
	TagID     string `bson:"tagID"`
}
