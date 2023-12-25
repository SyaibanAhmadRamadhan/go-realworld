package model

import (
	"time"
)

type Comment struct {
	ID        string    `bson:"id"`
	ArticleID string    `bson:"articleID"`
	AuthorID  string    `bson:"authorID"`
	Body      string    `bson:"body"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
