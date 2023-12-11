package model

type Article struct {
	ID          string `bson:"_id"`
	AuthorID    int    `bson:"authorID"`
	Slug        string `bson:"slug"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Body        string `bson:"body"`
	CreatedAt   string `bson:"createdAt"`
	UpdatedAt   string `bson:"updatedAt"`
}
