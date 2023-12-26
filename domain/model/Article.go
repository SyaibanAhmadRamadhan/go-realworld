package model

import (
	"time"
)

type Article struct {
	Id          string    `bson:"_id"          order:"true"`
	AuthorId    string    `bson:"authorId"     order:"false"`
	Slug        string    `bson:"slug"         order:"true"`
	Title       string    `bson:"title"        order:"true"`
	Description string    `bson:"description"  order:"false"`
	Body        string    `bson:"body"         order:"false"`
	CreatedAt   time.Time `bson:"createdAt"    order:"true"`
	UpdatedAt   time.Time `bson:"updatedAt"    order:"true"`
}
