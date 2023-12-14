package model

type ArticleTag struct {
	ArticleID string `bson:"articleID"`
	TagID     string `bson:"tagID"`
}

func (a *ArticleTag) TableName() string {
	return "articleTag"
}
