package model

import (
	"time"
)

type Article struct {
	ID          string    `bson:"_id"`
	AuthorID    string    `bson:"authorID"`
	Slug        string    `bson:"slug"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Body        string    `bson:"body"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}
