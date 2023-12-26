package model

import (
	"time"
)

type Comment struct {
	Id        string    `bson:"_id"`
	ArticleId string    `bson:"articleId"`
	AuthorId  string    `bson:"authorId"`
	Body      string    `bson:"body"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
