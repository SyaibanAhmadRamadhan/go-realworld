package model

type ArticleTag struct {
	ArticleID string   `bson:"articleID"`
	TagIDs    []string `bson:"tagIDs"`
}

func (a *ArticleTag) TableName() string {
	return "articleTag"
}
