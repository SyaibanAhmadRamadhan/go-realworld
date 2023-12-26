package model

const ArticleTagTableName string = "articleTag"

type ArticleTag struct {
	ArticleId string `bson:"articleId"`
	TagId     string `bson:"tagId"`
}
