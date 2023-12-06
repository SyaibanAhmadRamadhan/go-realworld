package model

type Comment struct {
	ID        int    `bson:"id"`
	ArticleID int    `bson:"articleID"`
	AuthorID  int    `bson:"authorID"`
	Body      string `bson:"body"`
	CreatedAt string `bson:"createdAt"`
	UpdatedAt string `bson:"updatedAt"`
}
